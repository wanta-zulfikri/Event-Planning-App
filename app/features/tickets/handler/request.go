package handler

type RequestCreateTicket struct {
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
}

type RequestUpdateTicket struct {
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
}
