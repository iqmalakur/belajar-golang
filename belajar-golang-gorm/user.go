package main

import "time"

// gorm automated mapping
// User => users
// OrderDetail => order_details
// ID => Primary Key
// or using gorm tag manually
// add TableName method to manually set the table name
type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        string    `gorm:"column:name"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdateddAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
}

func (u *User) TableName() string {
	return "users"
}
