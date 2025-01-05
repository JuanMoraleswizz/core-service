package repository

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm/clause"
	"uala.com/core-service/database"
	"uala.com/core-service/internal/entity"
)

type UserMariadbRepository struct {
	db database.Database
}

func NewUserMariadbRepository(db database.Database) *UserMariadbRepository {
	return &UserMariadbRepository{db: db}
}

func (r *UserMariadbRepository) Create(user *entity.User) error {
	log.Printf("InsertUserData: %v", user)
	result := r.db.GetDb().Clauses(clause.Returning{}).Select("id", "user_name").Create(user)
	if result.Error != nil {
		log.Errorf("InsertUserData: %v", result.Error)
		return result.Error
	}

	log.Printf("InsertUserData: %v", result.RowsAffected)
	return nil
}
