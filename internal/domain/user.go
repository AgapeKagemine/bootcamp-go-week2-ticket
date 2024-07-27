package domain

type User struct {
	ID       int      `json:"id,omitempty"`
	Username *string  `json:"username"`
	Phone    *string  `json:"phone"`
	Email    *string  `json:"email"`
	Balance  *float64 `json:"balance"`
}
