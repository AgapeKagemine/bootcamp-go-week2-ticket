package contract

import "net/http"

type Save interface {
	Save(w http.ResponseWriter, r *http.Request)
}
