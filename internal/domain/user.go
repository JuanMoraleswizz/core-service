package domain

import "time"

type User struct {
	ID       int64     `json:"id"`
	UserName string    `json:"user_name"`
	CreateAt time.Time `json:"created_at"`
}
