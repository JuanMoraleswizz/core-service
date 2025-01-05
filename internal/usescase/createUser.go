package usescase

import "uala.com/core-service/internal/domain"

type CreateUser interface {
	CreateUser(user domain.User) error
}
