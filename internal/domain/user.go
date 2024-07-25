package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required, notblank, alpha, len=2"`
	Phone    string `json:"phone" validate:"format=e164, required, notblank"`
	Email    string `json:"email" validate:"format=email, required, notblank"`
}
