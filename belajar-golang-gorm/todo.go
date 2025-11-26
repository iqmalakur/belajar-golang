package main

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserId      string
	Title       string
	Description string
}
