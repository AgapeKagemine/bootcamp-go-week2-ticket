// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package gotik

import (
	"database/sql"
)

type User struct {
	ID       int64          `json:"id"`
	Username sql.NullString `json:"username"`
	Phone    sql.NullString `json:"phone"`
	Email    sql.NullString `json:"email"`
	Balance  sql.NullString `json:"balance"`
}
