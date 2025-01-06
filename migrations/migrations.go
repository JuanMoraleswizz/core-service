package main

import (
	"fmt"

	"uala.com/core-service/config"
	"uala.com/core-service/database"
	"uala.com/core-service/internal/entity"
)

func main() {
	conf := config.GetConfig()
	db := database.NewMariadbDatabase(conf)
	retoMigrations(db)

}

func retoMigrations(db database.Database) {
	// Migrate the schema
	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE reto")
	db.GetDb().Exec(createDatabaseCommand)
	db.GetDb().Migrator().CreateTable(&entity.User{})
	db.GetDb().Migrator().CreateTable(&entity.Tweet{})
	db.GetDb().Migrator().CreateTable(&entity.Follow{})
}
