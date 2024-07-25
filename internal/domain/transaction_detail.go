package domain

type TransactionDetail struct {
	ID            int      `json:"id"`
	Time          string   `json:"time"`
	Status        string   `json:"status"`
	PaymentMethod string   `json:"payment_method"`
	TotalPayment  float64  `json:"total_payment"`
	User          User     `json:"user"`
	Event         Event    `json:"event"`
	Ticket        []Ticket `json:"ticket"`
}
