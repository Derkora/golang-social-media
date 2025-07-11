package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Bio      string    `json:"bio"`
}

type Post struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID  uuid.UUID `json:"user_id"`
	Content string    `json:"content"`
	Created time.Time `json:"created_at"`
}

type Like struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID uuid.UUID `json:"user_id"`
	PostID uuid.UUID `json:"post_id"`
}

type Comment struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID  uuid.UUID `json:"user_id"`
	PostID  uuid.UUID `json:"post_id"`
	Content string    `json:"content"`
}
