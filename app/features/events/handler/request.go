package handler

type RequestCreateEvent struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	EventDate   string `form:"date"`
	EventTime   string `form:"time"`
	Status      string `form:"status"`
	Category    string `form:"category"`
	Location    string `form:"location"`
	Image       string `form:"event_picture"`
}

type RequestUpdateEvent struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	EventDate   string `form:"date"`
	EventTime   string `form:"time"`
	Status      string `form:"status"`
	Category    string `form:"category"`
	Location    string `form:"location"`
	Image       string `form:"event_picture"`
}

// image masuk di update event
type RequestCreateEventWithTickets struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	EventDate   string                `json:"date"`
	EventTime   string                `json:"time"`
	Status      string                `json:"status"`
	Category    string                `json:"category"`
	Location    string                `json:"location"`
	Image       string                `json:"event_picture"`
	Tickets     []RequestCreateTicket `json:"tickets"`
}

type RequestCreateTicket struct {
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
}
