package handler

type RequestCreateTicket struct {
	TicketType     string `json:"tickettype"`
	TicketCategory string `json:"ticketcategory"`
	TicketPrice    uint   `json:"ticketprice"`
	TicketQuantity uint   `json:"ticketquantity"`
}

type RequestUpdateTicket struct {
	TicketType     string `json:"tickettype"`
	TicketCategory string `json:"ticketcategory"`
	TicketPrice    uint   `json:"ticketprice"`
	TicketQuantity uint   `json:"ticketquantity"`
}
