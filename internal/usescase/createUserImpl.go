package usescase

import (
	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/entity"
	"uala.com/core-service/internal/repository"
)

type CreateUserImpl struct {
	UserRepository repository.UserRepository
}

func NewCreateUserImpl(userRepository repository.UserRepository) CreateUser {
	return &CreateUserImpl{UserRepository: userRepository}
}

func (c *CreateUserImpl) CreateUser(user domain.User) error {
	userEntity := &entity.User{
		ID:       user.ID,
		UserName: user.UserName,
	}
	if err := c.UserRepository.Create(userEntity); err != nil {
		return err
	}
	return nil
}
