package entity

import "time"

type Follow struct {
	ID         int64     `gorm:"primary_key;auto_increment"`
	FollowerID int64     `gorm:"not null"`
	FolloweeID int64     `gorm:"not null"`
	CreateAt   time.Time `gorm:"not null;default:current_timestamp()"`
}
