package services

import (
	"errors"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
)

type EventService struct {
	r events.Repository
}

func New(r events.Repository) events.Service {
	return &EventService{r: r}
}

func (es *EventService) GetEventByEventID(id int) (events.Events, error) {
	event, err := es.r.GetEventByEventID(id)
	if err != nil {
		return events.Events{}, errors.New("terdapat masalah pada server")
	}
	return event, nil
}
