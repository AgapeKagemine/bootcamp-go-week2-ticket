package domain

type BuyTicket struct {
	TicketId *uint `json:"ticket_id"`
	Quantity *uint `json:"quantity"`
}

type EventBuyTicket struct {
	EventId *uint         `json:"event"`
	UserId  *uint         `json:"user"`
	Ticket  *[]*BuyTicket `json:"ticket"`
}
