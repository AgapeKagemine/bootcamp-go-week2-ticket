package domain

type Event struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Date        string   `json:"date"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Ticket      []Ticket `json:"ticket"`
}
