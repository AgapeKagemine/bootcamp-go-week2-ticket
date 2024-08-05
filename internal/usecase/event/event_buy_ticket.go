package event

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gotik/internal/domain"
)

type BuyTicket interface {
	BuyTicket(context.Context) (domain.TransactionDetail, error)
}

func (uc *EventUsecaseImpl) BuyTicket(ctx context.Context) (h domain.TransactionDetail, err error) {
	mtx := &sync.Mutex{}
	mtx.Lock()
	defer mtx.Unlock()

	history := &domain.TransactionDetail{
		Time: ctx.Value(domain.Start("start")).(time.Time).String(),
	}

	request := ctx.Value(domain.Start("request")).(*domain.EventBuyTicket)

	if request == nil {
		return *history, fmt.Errorf("invalid request")
	}

	history.Time = ctx.Value(domain.Start("start")).(time.Time).String()

	user, err := uc.userRepo.FindById(ctx, int(*request.UserId))
	if err != nil || user.ID == 0 {
		history.Status = "Failed"
		return *history, fmt.Errorf("user with id: %d not found", *request.UserId)
	}

	event, err := uc.eventRepo.FindById(ctx, int(*request.EventId))
	if err != nil || event.ID == 0 {
		history.Status = "Failed"
		return *history, fmt.Errorf("event with id: %d not found", *request.EventId)
	}

	defer func() (domain.TransactionDetail, error) {
		if err != nil {
			history.Status = "Failed"
		}
		h, _ := uc.tdRepo.Save(ctx, history)
		err = uc.tdRepo.SaveTransactionDetailsEventsUsers(ctx, history.ID, event.ID, user.ID)
		h.ID = history.ID
		return h, err
	}()

	history.Event = event
	history.Event.Ticket = []domain.Ticket{}

	for _, req := range *request.Ticket {
		for _, ticket := range event.Ticket {
			if ticket.ID == int(*req.TicketId) {
				history.TotalPayment += ticket.Price * float64(*req.Quantity)
				break
			}
		}
	}

	if history.TotalPayment > *user.Balance {
		return *history, fmt.Errorf("user %s with id: %d balance(s) is not enough", *user.Username, user.ID)
	}

	*user.Balance -= history.TotalPayment
	history.User = user

	for t, req := range *request.Ticket {
		for i, ticket := range event.Ticket {
			if ticket.Stock < *req.Quantity {
				return *history, fmt.Errorf("%s ticket(s) left: %d", ticket.Type, ticket.Stock)
			}

			if ticket.ID == int(*req.TicketId) {
				event.Ticket[i].Stock = event.Ticket[i].Stock - *req.Quantity
				history.Event.Ticket = append(history.Event.Ticket, ticket)
				history.Event.Ticket[t].Stock = *req.Quantity
				break
			}

			if i == len(event.Ticket)-1 {
				return *history, fmt.Errorf("ticket with %d not found", *req.TicketId)
			}

		}
	}

	_, err = uc.userRepo.Update(ctx, &user)
	if err != nil {
		return *history, err
	}

	_, err = uc.eventRepo.Update(ctx, &event)
	if err != nil {
		return *history, err
	}

	for _, t := range event.Ticket {
		_, err := uc.ticketRepo.Update(ctx, &t)
		if err != nil {
			return *history, err
		}
	}

	history.Status = "Success"
	return *history, nil
}
