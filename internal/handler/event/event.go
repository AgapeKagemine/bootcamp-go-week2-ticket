package event

import (
	"context"
	"gotik/internal/domain"
	"gotik/internal/handler/contract"
	"gotik/internal/usecase/event"
)

type EventHandler interface {
	contract.FindAll
	contract.FindById
	contract.Save
	contract.Update
	contract.DeleteById
	BuyTicket
}

type EventHandlerImpl struct {
	uc event.EventUsecase[context.Context, domain.Event]
}

func NewEventHandler(uc event.EventUsecase[context.Context, domain.Event]) EventHandler {
	return &EventHandlerImpl{
		uc: uc,
	}
}
