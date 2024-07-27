package ticket

import (
	contract "gotik/internal/handler/contract/http"
	"gotik/internal/usecase/ticket"
)

type TicketHandler interface {
	contract.FindAll
	contract.FindById
	contract.Save
}

type TicketHandlerImpl struct {
	uc ticket.TicketUsecase
}

func NewTicketHandler(uc ticket.TicketUsecase) TicketHandler {
	return &TicketHandlerImpl{
		uc: uc,
	}
}
