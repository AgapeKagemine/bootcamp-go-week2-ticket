package event

import (
	"gotik/internal/handler/contract"
	"gotik/internal/usecase/event"
)

type EventHandler interface {
	contract.FindAll
	contract.FindById
	contract.Save
	contract.Update
	contract.DeleteById
}

type EventHandlerImpl struct {
	uc event.EventUsecase
}

func NewEventHandler(uc event.EventUsecase) EventHandler {
	return EventHandlerImpl{
		uc: uc,
	}
}
