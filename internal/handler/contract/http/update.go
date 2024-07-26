package contract

import "net/http"

type Update interface {
	Update(w http.ResponseWriter, r *http.Request)
}
