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
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdateddAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
}

func (u *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

type UserLog struct {
	ID        string `gorm:"autoIncrement"`
	UserId    string
	Action    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
