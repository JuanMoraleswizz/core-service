package domain

import "time"

type Follow struct {
	ID         int64     `json:"id"`
	FollowerID int64     `json:"follower_id"`
	FolloweeID int64     `json:"Followee_id"`
	CreateAt   time.Time `json:"created_at"`
}

