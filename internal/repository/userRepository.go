package repository

import "uala.com/core-service/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
}
