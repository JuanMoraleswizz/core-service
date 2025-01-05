package repository

import (
	"strconv"

	"github.com/labstack/gommon/log"
	"uala.com/core-service/database"
	"uala.com/core-service/internal/entity"
	"uala.com/core-service/internal/rabbit"
)

type FollowMariadbRepository struct {
	db     database.Database
	rabbit *rabbit.Rabbit
}

func NewFollowMariadbRepository(db database.Database, rabbit *rabbit.Rabbit) *FollowMariadbRepository {
	return &FollowMariadbRepository{db: db, rabbit: rabbit}
}

func (r *FollowMariadbRepository) Follow(follow entity.Follow) error {

	result := r.db.GetDb().Create(follow)
	if result.Error != nil {
		log.Errorf("InsertFollowData: %v", result.Error)
		return result.Error
	}
	mainChanel, err := r.rabbit.GetChannel()
	if err != nil {
		log.Errorf("Error obteniendo el chanel: %v", err)
	}
	mainQueue, err := r.rabbit.CreateQueue(mainChanel)
	if err != nil {
		log.Errorf("Error declarando la queue: %v", err)
	}
	var userString string
	userString = strconv.FormatInt(follow.FollowerID, 10)
	err = r.rabbit.Publish(mainChanel, mainQueue, []byte(userString))
	if err != nil {
		log.Errorf("Error publicando en la queue: %v", err)
	}

	log.Printf("InsertFollowData: %v", result.RowsAffected)
	return nil
}

func (r *FollowMariadbRepository) Unfollow(follow entity.Follow) error {

	result := r.db.GetDb().Delete(follow)
	if result.Error != nil {
		log.Errorf("DeleteFollowData: %v", result.Error)
		return result.Error
	}
	mainChanel, err := r.rabbit.GetChannel()
	if err != nil {
		log.Errorf("Error obteniendo el chanel: %v", err)
	}
	mainQueue, err := r.rabbit.CreateQueue(mainChanel)
	if err != nil {
		log.Errorf("Error declarando la queue: %v", err)
	}
	var userString string
	userString = strconv.FormatInt(follow.FollowerID, 10)
	err = r.rabbit.Publish(mainChanel, mainQueue, []byte(userString))
	if err != nil {
		log.Errorf("Error publicando en la queue: %v", err)
	}

	log.Printf("DeleteFollowData: %v", result.RowsAffected)
	return nil
}
