package main

import "time"

type Wallet struct {
	ID        string
	UserId    string
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
