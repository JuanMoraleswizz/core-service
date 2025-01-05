package entity

import "time"

type User struct {
	ID       int64     `gorm:"primary_key;auto_increment"`
	UserName string    `gorm:"not null;unique_index"`
	CreateAt time.Time `gorm:"not null;default:now()"`
}
