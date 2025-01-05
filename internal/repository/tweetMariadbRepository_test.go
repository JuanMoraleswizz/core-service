package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"uala.com/core-service/internal/entity"
)

// Mock para RabbitMQ Channel
type MockChannel struct {
	mock.Mock
}

func (m *MockChannel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp091.Table) (amqp091.Queue, error) {
	args := m.Called(name, durable, autoDelete, exclusive, noWait, args)
	return args.Get(0).(amqp091.Queue), args.Error(1)
}

func (m *MockChannel) PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp091.Publishing) error {
	args := m.Called(ctx, exchange, key, mandatory, immediate, msg)
	return args.Error(0)
}

func (m *MockChannel) Close() error {
	args := m.Called()
	return args.Error(0)
}

// Mock para Rabbit que implementa la interfaz rabbit.Rabbit
type MockRabbit struct {
	mock.Mock
}

func (m *MockRabbit) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRabbit) GetChannel() (rabbit.Channel, error) {
	args := m.Called()
	if ch, ok := args.Get(0).(rabbit.Channel); ok {
		return ch, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestTweetMariadbRepository_Create(t *testing.T) {
	tests := []struct {
		name            string
		tweet           *entity.Tweet
		setupDBMock     func(sqlmock.Sqlmock)
		setupRabbitMock func(*MockChannel)
		expectedError   bool
	}{
		{
			name: "successful creation",
			tweet: &entity.Tweet{
				UserID:  123,
				Content: "Test tweet",
			},
			setupDBMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `tweets`").
					WithArgs(123, "Test tweet"). // ajusta según los campos de tu entidad Tweet
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			setupRabbitMock: func(mockChannel *MockChannel) {
				mockChannel.On("QueueDeclare",
					"main", true, false, false, false, mock.Anything).
					Return(amqp091.Queue{Name: "main"}, nil)

				mockChannel.On("PublishWithContext",
					mock.Anything, "", "main", false, false, mock.Anything).
					Return(nil)

				mockChannel.On("Close").Return(nil)
			},
			expectedError: false,
		},
		{
			name: "database error",
			tweet: &entity.Tweet{
				UserID:  123,
				Content: "Test tweet",
			},
			setupDBMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `tweets`").
					WithArgs(123, "Test tweet").
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			setupRabbitMock: func(mockChannel *MockChannel) {
				// No se espera que se llame a RabbitMQ si hay error en DB
			},
			expectedError: true,
		},
		{
			name: "rabbit error",
			tweet: &entity.Tweet{
				UserID:  123,
				Content: "Test tweet",
			},
			setupDBMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `tweets`").
					WithArgs(123, "Test tweet").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			setupRabbitMock: func(mockChannel *MockChannel) {
				mockChannel.On("QueueDeclare",
					"main", true, false, false, false, mock.Anything).
					Return(amqp091.Queue{}, errors.New("rabbit error"))

				mockChannel.On("Close").Return(nil)
			},
			expectedError: false, // según tu implementación actual, los errores de rabbit no causan error en Create
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sqlDB, dbMock, mockDB := newMockDB(t)
			defer sqlDB.Close()

			mockChannel := new(MockChannel)
			mockRabbit := new(MockRabbit)
			mockRabbit.On("GetChannel").Return(mockChannel, nil)

			tt.setupDBMock(dbMock)
			tt.setupRabbitMock(mockChannel)

			repo := NewTweetMariadbRepository(mockDB, mockRabbit)

			// Act
			err := repo.Create(tt.tweet)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify expectations
			if err := dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("hay expectativas de base de datos sin cumplir: %s", err)
			}
			mockChannel.AssertExpectations(t)
			mockRabbit.AssertExpectations(t)
		})
	}
}
