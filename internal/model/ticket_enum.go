package model

type TicketStatus int8

const (
	StatusActive TicketStatus = iota
	StatusEaten
)

func StringToTicketStatus(s string) TicketStatus {
	switch s {
	case "active":
		return StatusActive
	case "eaten":
		return StatusEaten
	default:
		return StatusActive
	}
}

func (s TicketStatus) ToString() string {
	switch s {
	case StatusActive:
		return "active"
	case StatusEaten:
		return "eaten"
	default:
		return "active"
	}
}
