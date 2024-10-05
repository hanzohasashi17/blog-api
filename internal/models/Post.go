package models

import (
	"time"
)

type Post struct {
	Id        int       `json:"id"`
	Title     string    `json:"title" validate:"required,min=2,max=100"`
	Content   string    `json:"content" validate:"required,min=2,max=1000"`
	Author    string    `json:"author" validate:"required,min=2,max=50"`
	CreatedAt time.Time `json:"created_at"`
}
