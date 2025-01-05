package repository

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"uala.com/core-service/database"
	"uala.com/core-service/internal/entity"
	"uala.com/core-service/internal/rabbit"
)

type TweetMariadbRepository struct {
	db     database.Database
	rabbit *rabbit.Rabbit
}

func NewTweetMariadbRepository(db database.Database, rabbit rabbit.Rabbit) *TweetMariadbRepository {
	return &TweetMariadbRepository{db: db, rabbit: &rabbit}
}

func (r *TweetMariadbRepository) Create(tweet *entity.Tweet) error {
	log.Printf("InsertTweetData: %v", tweet)
	result := r.db.GetDb().Create(tweet)
	if result.Error != nil {
		log.Errorf("InsertTweetData: %v", result.Error)
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
	userString = fmt.Sprintf("%d", tweet.UserID)
	err = r.rabbit.Publish(mainChanel, mainQueue, []byte(userString))

	var follow []entity.Follow
	r.db.GetDb().Where("followee_id = ?", tweet.UserID).Find(&follow)
	fmt.Println(follow)
	fmt.Println("realizo busqueda")
	for _, f := range follow {
		fmt.Println(f.FollowerID)
		userString = fmt.Sprintf("%d", f.FollowerID)
		err = r.rabbit.Publish(mainChanel, mainQueue, []byte(userString))
		if err != nil {
			log.Errorf("Error publicando en la queue: %v", err)
		}
	}

	defer mainChanel.Close()
	log.Printf("InsertTweetData: %v", result.RowsAffected)
	return nil
}
