package entity

import "time"

type Tweet struct {
	ID       int64     `gorm:"primary_key;auto_increment"`
	Content  string    `gorm:"not null"`
	UserID   int64     `gorm:"not null"`
	CreateAt time.Time `gorm:"not null;default:now()"`
}
