package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"uala.com/core-service/internal/entity"
)

type MockDatabase struct {
	DB *gorm.DB
}

func (m *MockDatabase) GetDb() *gorm.DB {
	return m.DB
}

func newMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *MockDatabase) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening gorm connection: %v", err)
	}

	return sqlDB, mock, &MockDatabase{DB: db}
}

func TestUserMariadbRepository_Create(t *testing.T) {
	tests := []struct {
		name          string
		user          *entity.User
		setupMock     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name: "successful creation",
			user: &entity.User{
				UserName: "testuser",
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `users`").
					WithArgs("testuser"). // ajusta los argumentos según los campos de tu entidad User
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name: "database error",
			user: &entity.User{
				UserName: "testuser",
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `users`").
					WithArgs("testuser").
					WillReturnError(sql.ErrConnDone)
				mock.ExpectRollback()
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sqlDB, mock, mockDB := newMockDB(t)
			defer sqlDB.Close()

			tt.setupMock(mock)
			repo := NewUserMariadbRepository(mockDB)

			// Act
			err := repo.Create(tt.user)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verificar que todas las expectativas del mock se cumplieron
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("hay expectativas sin cumplir: %s", err)
			}
		})
	}
}
