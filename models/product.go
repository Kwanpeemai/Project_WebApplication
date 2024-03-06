package models

import "time"

type Yogurt struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Yogurt) TableName() string {
	return "Posts"
}
