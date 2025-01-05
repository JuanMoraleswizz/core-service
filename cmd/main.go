package main

import (
	"uala.com/core-service/cmd/server"
	"uala.com/core-service/config"
	"uala.com/core-service/database"
	"uala.com/core-service/internal/rabbit"
)

func main() {
	conf := config.GetConfig()
	db := database.NewMariadbDatabase(conf)
	rabbit := rabbit.NewRabbit(conf)
	server.NewServer(conf, db, *rabbit).Start()
}
