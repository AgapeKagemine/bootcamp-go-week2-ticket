package event

import (
	contract "gotik/internal/handler/contract/http"
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
	uc event.EventUsecase
}

func NewEventHandler(uc event.EventUsecase) EventHandler {
	return &EventHandlerImpl{
		uc: uc,
	}
}
