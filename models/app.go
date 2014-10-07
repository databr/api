package models

import "time"

type App struct {
	Name      string
	Token     string
	UserId    int64
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
