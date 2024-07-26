package event

import (
	"gotik/internal/usecase/event"
)

type EventHandler interface {
	event.EventUsecase
}

type EventHandlerImpl struct {
	uc event.EventUsecase
}

func NewEventHandler(uc event.EventUsecase) EventHandler {
	return EventHandlerImpl{
		uc: uc,
	}
}
