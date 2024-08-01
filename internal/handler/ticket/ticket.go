package ticket

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/handler/contract"
	"gotik/internal/usecase/ticket"
)

type TicketHandler interface {
	contract.FindAll
	contract.FindById
	contract.Save
}

type TicketHandlerImpl struct {
	uc ticket.TicketUsecase[context.Context, domain.Ticket]
}

func NewTicketHandler(uc ticket.TicketUsecase[context.Context, domain.Ticket]) TicketHandler {
	return &TicketHandlerImpl{
		uc: uc,
	}
}
