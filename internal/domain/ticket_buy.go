package domain

type EventBuyTicket struct {
	EventId  *uint   `json:"event"`
	UserId   *uint   `json:"user"`
	TicketId *[]uint `json:"ticket"`
}
