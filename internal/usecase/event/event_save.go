package event

import (
	"context"
	"time"

	"gotik/internal/domain"
)

// Save implements EventUsecase.
func (uc *EventUsecaseImpl) Save(ct context.Context, e *domain.Event) (event domain.Event, err error) {
	for _, ev := range *populateEvent() {
		event, err = uc.eventRepo.Save(ct, &ev)
		if err != nil {
			return domain.Event{}, err
		}

		ctx := context.WithValue(ct, domain.Start("eventID"), &event.ID)

		for _, t := range ev.Ticket {
			ticket, err := uc.ticketRepo.Save(ctx, &t)
			if err != nil {
				return domain.Event{}, err
			}

			err = uc.ticketRepo.SaveEventTicket(ctx, ticket.ID)
			if err != nil {
				return domain.Event{}, err
			}
		}
	}
	return domain.Event{}, nil
}

func populateEvent() *[]domain.Event {
	return &[]domain.Event{
		{
			Name:        "Event 1",
			Date:        time.Date(2022, time.February, 14, 15, 30, 0, 0, time.Local).String(),
			Description: "Description 1",
			Location:    "Location 1",
			Ticket: []domain.Ticket{
				{
					Stock: 10,
					Type:  "VIP",
					Price: 5000,
				},
				{
					Stock: 100,
					Type:  "CAT 1",
					Price: 250,
				},
			},
		},
		{
			Name:        "Event 2",
			Date:        time.Date(2022, time.February, 14, 15, 30, 0, 0, time.Local).String(),
			Description: "Description 2",
			Location:    "Location 2",
			Ticket: []domain.Ticket{
				{
					Stock: 10,
					Type:  "VIP",
					Price: 5000,
				},
				{
					Stock: 100,
					Type:  "CAT 1",
					Price: 250,
				},
			},
		},
	}
}
