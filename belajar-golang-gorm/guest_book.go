package main

import "time"

type GuestBook struct {
	ID        int64 `gorm:"autoIncrement"`
	Name      string
	Email     string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
