package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"uala.com/core-service/config"
	"uala.com/core-service/database"
	"uala.com/core-service/internal/handlers"
	"uala.com/core-service/internal/rabbit"
	"uala.com/core-service/internal/repository"
	"uala.com/core-service/internal/usescase"
)

type server struct {
	app    *echo.Echo
	db     database.Database
	conf   *config.Config
	rabbit rabbit.Rabbit
}

func NewServer(conf *config.Config, db database.Database, rabbit rabbit.Rabbit) *server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	return &server{
		app:    echoApp,
		db:     db,
		conf:   conf,
		rabbit: rabbit,
	}
}

func (s *server) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	// Health check adding
	s.app.GET("/ping", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	s.initializeCockroachHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *server) initializeCockroachHttpHandler() {
	// Initialize all repositories
	userRepository := repository.NewUserMariadbRepository(s.db)
	followRepository := repository.NewFollowMariadbRepository(s.db)
	tweetRepository := repository.NewTweetMariadbRepository(s.db, s.rabbit)

	// Initialize all usecases
	createUserUseCase := usescase.NewCreateUserImpl(userRepository)
	createTweetUseCase := usescase.NewCreateTweetImpl(tweetRepository)
	createFollowUseCase := usescase.NewFollowUserImpl(followRepository)

	// Initialize all handlers
	userHandler := handlers.NewUserHandler(createUserUseCase)
	tweetHandler := handlers.NewTweetHandler(createTweetUseCase)
	followHandler := handlers.NewFollowHandler(createFollowUseCase)

	// Routers
	s.app.POST("/v1/user", userHandler.CreateUser)
	s.app.POST("/v1/tweet", tweetHandler.CreateTweet)
	s.app.POST("/v1/follow", followHandler.FollowUser)
	s.app.DELETE("/v1/follow", followHandler.UnfollowUser)

}
