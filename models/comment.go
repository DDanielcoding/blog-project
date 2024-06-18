package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	BlogID    uint      `json:"blog_id"`
	AuthorID  uint      `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
