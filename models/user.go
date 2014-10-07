package models

import "time"

type User struct {
	Id        int64
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
