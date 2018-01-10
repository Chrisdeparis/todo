package model

import "time"

type Task struct {
	ID        string    `json:"id,omitempty" bson:"_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Status    bool      `json:"status"`
}
