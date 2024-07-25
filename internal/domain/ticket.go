package domain

type TicketType struct {
	Type  string  `json:"type"`
	Price float64 `json:"price" validate:"gte=0, required"`
}

type Ticket struct {
	ID    int        `json:"id"`
	Stock uint       `json:"stock" validate:"gte=0, required"`
	Type  TicketType `json:"type" validate:"dive"`
}
