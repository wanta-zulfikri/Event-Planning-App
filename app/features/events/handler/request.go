package handler

type RequestCreateEvent struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	EventDate   string `form:"eventdate"`
	EventTime   string `form:"eventtime"`
	Status      string `form:"status"`
	Category    string `form:"category"`
	Location    string `form:"location"`
	Image       string `form:"image"`
	Username    string `form:"username"`
}

type RequestUpdateEvent struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	EventDate   string `form:"eventdate"`
	EventTime   string `form:"eventtime"`
	Status      string `form:"status"`
	Category    string `form:"category"`
	Location    string `form:"location"`
	Image       string `form:"image"`
	Username    string `form:"username"`
}

type RequestCreateEventWithTickets struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	EventDate   string                `json:"eventdate"`
	EventTime   string                `json:"eventtime"`
	Status      string                `json:"status"`
	Category    string                `json:"category"`
	Location    string                `json:"location"`
	Image       string                `json:"image"`
	Tickets     []RequestCreateTicket `json:"tickets"`
}

type RequestCreateTicket struct {
	Title          string `json:"title"`
	TicketType     string `json:"ticket_type"`
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
}
