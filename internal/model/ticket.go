package model

import (
	"ticket/internal/helpers"
	"time"
)

type Ticket struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Number string
	Status TicketStatus

	UserID uint
}

func NewTicket(userId uint) *Ticket {

	return &Ticket{
		Number: helpers.GenerateTicketNumber(),
		Status: StatusActive,
		UserID: userId,
	}
}
