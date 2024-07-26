package contract

import "net/http"

type DeleteById interface {
	DeleteById(w http.ResponseWriter, r *http.Request)
}
