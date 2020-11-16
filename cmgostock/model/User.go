package model

import "time"

// https://gorm.io/docs/models.html

type User struct {
	ID        uint
	Username  string
	Password  string
	Level     string
	CreatedAt time.Time
}
