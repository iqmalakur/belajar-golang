package main

import "time"

type Product struct {
	ID           string
	Name         string
	Price        int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LikedByUsers []User `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id"`
}
