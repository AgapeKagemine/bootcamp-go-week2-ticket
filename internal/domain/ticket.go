package domain

type Ticket struct {
	ID    int     `json:"id"`
	Stock uint    `json:"stock"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
