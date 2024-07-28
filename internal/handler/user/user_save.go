package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gotik/internal/domain"

	"github.com/rs/zerolog/log"
)

// Save implements UserHandler.
func (h *UserHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	c := context.WithValue(r.Context(), domain.Start("start"), time.Now().Local())
	ctx, cancel := context.WithDeadline(c, time.Now().Local().Add(time.Second*30))

	w.Header().Set("Content-Type", "application/json")

	response := &domain.ResponseBody{
		StatusCode: 0,
		Message:    "",
		Payload:    nil,
	}

	defer func() {
		w.WriteHeader(int(response.StatusCode))
		json.NewEncoder(w).Encode(response)
		r.Body.Close()
		cancel()
		log.Info().Uint("httpStatus", response.StatusCode).Str("statusDesc", response.Message).Str("processTime", time.Now().Local().Sub(ctx.Value(domain.Start("start")).(time.Time)).String()).Msg(fmt.Sprintf("USER SAVE - %s", http.StatusText(int(response.StatusCode))))
	}()

	if r.Method != http.MethodPost {
		response.StatusCode = http.StatusMethodNotAllowed
		response.Message = http.StatusText(http.StatusMethodNotAllowed)
		return
	}

	if len(r.Header.Values("Content-Type")) == 0 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Bad request - Content-type is not set"
		return
	}

	if r.Header.Values("Content-Type")[0] != "application/json" {
		response.StatusCode = http.StatusUnsupportedMediaType
		response.Message = http.StatusText(http.StatusUnsupportedMediaType)
		return
	}

	response.StatusCode = http.StatusBadRequest

	if r.Body == nil {
		response.Message = "Bad request - Body is empty"
		return
	}

	var request domain.User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = "Bad request - Bad JSON format"
		return
	}

	defer func() { response.Payload = request }()

	if request.Username == nil || *request.Username == "" || len(*request.Username) < 2 {
		response.Message = "Username is required and/or less than 2 characters"
		return
	}

	if request.Phone == nil || *request.Phone == "" {
		response.Message = "Phone is required"
		return
	}

	if request.Email == nil || *request.Email == "" {
		response.Message = "Email is required"
		return
	}

	if request.Balance == nil || *request.Balance < 0 {
		response.Message = "Balance is required and/or cannot be negative"
		return
	}

	err = h.uc.Save(ctx, &request)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}

	response.StatusCode = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
}
