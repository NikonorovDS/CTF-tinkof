package model

import (
	"ticket/internal/helpers"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Username string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100)"`

	Balance uint `gorm:"default:5000"`

	Luck uint `gorm:"default:0"`

	Tickets []Ticket
}

func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: helpers.HashPassword(password),
	}
}
