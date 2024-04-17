package storage

type Store interface {
	Tickets() TicketRepository
	Users() UserRepository
	Ping() error
}
