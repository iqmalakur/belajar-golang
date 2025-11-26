package main

import "time"

type Address struct {
	ID        int64 `gorm:"autoIncrement"`
	UserId    string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:user_id;references:id"`
}
