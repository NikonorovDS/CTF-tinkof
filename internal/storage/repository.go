package storage

import "ticket/internal/model"

type TicketRepository interface {
	BuyTicket(userId uint) (*model.Ticket, error)
	EatTicket(userId, ticketId uint) (int, string, error)
	AllByUserId(userId uint) ([]*model.Ticket, error)
	FindById(id uint) (*model.Ticket, error)
	Create(ticket *model.Ticket) (*model.Ticket, error)
	Update(ticket *model.Ticket) (*model.Ticket, error)
	Delete(id uint) error
	Migrate() error
}

type UserRepository interface {
	FindById(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id uint) error
	Migrate() error
}
