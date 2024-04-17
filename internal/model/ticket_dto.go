package model

type TicketDTO struct {
	ID     uint   `json:"id"`
	Number string `json:"number"`
	Status string `json:"status"`
}

type TicketReqDTO struct {
	ID uint `json:"id"`
}

func (dto *TicketDTO) ToModel() *Ticket {
	return &Ticket{
		ID:     dto.ID,
		Number: dto.Number,
		Status: StringToTicketStatus(dto.Status),
	}
}

func (m *Ticket) ToDTO() *TicketDTO {
	return &TicketDTO{
		ID:     m.ID,
		Number: m.Number,
		Status: m.Status.ToString(),
	}
}

func ToTicketDTOs(tickets []*Ticket) []*TicketDTO {
	ticketDTOs := make([]*TicketDTO, len(tickets))

	for i, ticket := range tickets {
		ticketDTOs[i] = ticket.ToDTO()
	}

	return ticketDTOs
}
