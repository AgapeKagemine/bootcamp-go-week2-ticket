package event

import (
	"context"
	"fmt"
	"gotik/internal/domain"
	"sync"
	"time"
)

type BuyTicket interface {
	BuyTicket(context.Context) (domain.TransactionDetail, error)
}

func (uc *EventUsecaseImpl) BuyTicket(ctx context.Context) (history domain.TransactionDetail, err error) {
	mtx := &sync.Mutex{}
	mtx.Lock()
	defer mtx.Unlock()
	request := ctx.Value(domain.Start("request")).(*domain.EventBuyTicket)

	history = domain.TransactionDetail{
		Time:         ctx.Value(domain.Start("start")).(time.Time).String(),
		Status:       "Pending",
		TotalPayment: 0,
		User:         domain.User{},
		Event:        domain.Event{},
	}

	defer func() (domain.TransactionDetail, error) {
		if err != nil {
			history.Status = "Failed"
		}
		history, err := uc.tdRepo.Save(ctx, &history)
		return history, err
	}()

	user, err := uc.userRepo.FindById(ctx, int(*request.UserId))
	if err != nil {
		return history, err
	}
	history.User = user

	event, err := uc.eventRepo.FindById(ctx, int(*request.EventId))
	if err != nil {
		return history, err
	}

	history.Event = event
	history.Event.Ticket = []domain.Ticket{}

	for _, req := range *request.Ticket {
		for _, ticket := range event.Ticket {
			if ticket.ID == int(*req.TicketId) {
				history.TotalPayment += ticket.Type.Price * float64(*req.Quantity)
				break
			}
		}
	}

	if history.TotalPayment > *user.Balance {
		return history, fmt.Errorf("user %s with id: %d balance(s) is not enough", *user.Username, user.ID)
	}

	*user.Balance -= history.TotalPayment
	history.User = user

	for r, req := range *request.Ticket {
		for i, ticket := range event.Ticket {
			if ticket.ID == int(*req.TicketId) {
				if ticket.Stock < *req.TicketId {
					return history, fmt.Errorf("%s ticket(s) left: %d", ticket.Type.Type, ticket.Stock)
				}

				event.Ticket[i].Stock -= *req.Quantity
				history.Event.Ticket = append(history.Event.Ticket, ticket)
				history.Event.Ticket[r].Stock = *req.Quantity
				break
			}

			if i == len(event.Ticket)-1 {
				return history, fmt.Errorf("ticket with %d not found", *req.TicketId)
			}

		}
	}

	user, err = uc.userRepo.Update(ctx, &user)
	if err != nil {
		return history, err
	}

	event, err = uc.eventRepo.Update(ctx, &event)
	if err != nil {
		return history, err
	}

	history.Status = "Success"
	return
}
