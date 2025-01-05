package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"uala.com/core-service/config"
)

type mariadbDatabase struct {
	db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *mariadbDatabase
)

func NewMariadbDatabase(config *config.Config) *mariadbDatabase {
	once.Do(func() {
		dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.DBName, config.Db.Charset)
		db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		fmt.Println(dns)
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Println("Connected to database")
		dbInstance = &mariadbDatabase{db: db}

	})
	return dbInstance
}

func (mb *mariadbDatabase) GetDb() *gorm.DB {
	return mb.db
}
