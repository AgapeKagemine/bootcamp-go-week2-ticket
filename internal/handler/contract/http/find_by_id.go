package contract

import "net/http"

type FindById interface {
	FindById(w http.ResponseWriter, r *http.Request)
}
