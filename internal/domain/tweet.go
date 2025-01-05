package domain

import "time"

type Tweet struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
