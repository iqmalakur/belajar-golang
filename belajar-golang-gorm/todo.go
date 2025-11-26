package main

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          int64 `gorm:"autoIncrement"`
	UserId      string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
