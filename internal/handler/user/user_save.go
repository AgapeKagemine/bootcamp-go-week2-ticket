package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gotik/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Save implements UserHandler.
func (h *UserHandlerImpl) Save(c *gin.Context) {
	ct := context.WithValue(c.Request.Context(), domain.Start("start"), time.Now().Local())
	ctx, cancel := context.WithDeadline(ct, time.Now().Local().Add(time.Second*30))

	c.Writer.Header().Set("Content-Type", "application/json")

	response := &domain.ResponseBody{
		StatusCode: 0,
		Message:    "",
		Payload:    nil,
	}

	defer func() {
		c.JSON(int(response.StatusCode), response)
		c.Request.Body.Close()
		cancel()
		log.Info().Uint("httpStatus", response.StatusCode).Str("statusDesc", response.Message).Str("processTime", time.Now().Local().Sub(ctx.Value(domain.Start("start")).(time.Time)).String()).Msg(fmt.Sprintf("USER SAVE - %s", http.StatusText(int(response.StatusCode))))
	}()

	if c.Request.Method != http.MethodPost {
		response.StatusCode = http.StatusMethodNotAllowed
		response.Message = http.StatusText(http.StatusMethodNotAllowed)
		return
	}

	if len(c.Request.Header.Values("Content-Type")) == 0 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Bad request - Content-type is not set"
		return
	}

	if c.Request.Header.Values("Content-Type")[0] != "application/json" {
		response.StatusCode = http.StatusUnsupportedMediaType
		response.Message = http.StatusText(http.StatusUnsupportedMediaType)
		return
	}

	response.StatusCode = http.StatusBadRequest

	if c.Request.Body == nil {
		response.Message = "Bad request - Body is empty"
		return
	}

	var request domain.User
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		response.Message = "Bad request - Bad JSON format"
		return
	}

	response.Payload = request

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

	event, err := h.uc.Save(ctx, &request)
	response.Payload = event

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}

	response.StatusCode = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
}
