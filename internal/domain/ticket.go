package domain

type TicketType struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type Ticket struct {
	ID    int        `json:"id"`
	Stock uint       `json:"stock"`
	Type  TicketType `json:"type"`
}
