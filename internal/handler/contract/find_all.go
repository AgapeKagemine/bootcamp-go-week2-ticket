package contract

import "net/http"

type FindAll interface {
	FindAll(w http.ResponseWriter, r *http.Request)
}
