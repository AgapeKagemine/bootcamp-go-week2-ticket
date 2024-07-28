package event

import (
	"context"
	"time"

	"gotik/internal/domain"
)

// Save implements EventUsecase.
func (uc *EventUsecaseImpl) Save(ctx context.Context, event *domain.Event) (err error) {
	for _, e := range *populateEvent() {
		err = uc.eventRepo.Save(ctx, &e)
		if err != nil {
			return err
		}
	}
	return nil
}

func populateEvent() *[]domain.Event {
	return &[]domain.Event{
		{
			ID:          1,
			Name:        "Event 1",
			Date:        time.Date(2022, time.February, 14, 15, 30, 0, 0, time.Local).String(),
			Description: "Description 1",
			Location:    "Location 1",
			Ticket: []domain.Ticket{
				{
					ID:    1,
					Stock: 10,
					Type: domain.TicketType{
						Type:  "VIP",
						Price: 5000,
					},
				},
				{
					ID:    2,
					Stock: 100,
					Type: domain.TicketType{
						Type:  "CAT 1",
						Price: 250,
					},
				},
			},
		},
		{
			ID:          2,
			Name:        "Event 2",
			Date:        time.Date(2022, time.February, 14, 15, 30, 0, 0, time.Local).String(),
			Description: "Description 2",
			Location:    "Location 2",
			Ticket: []domain.Ticket{
				{
					ID:    1,
					Stock: 10,
					Type: domain.TicketType{
						Type:  "VIP",
						Price: 5000,
					},
				},
				{
					ID:    2,
					Stock: 100,
					Type: domain.TicketType{
						Type:  "CAT 1",
						Price: 250,
					},
				},
			},
		},
	}
}
