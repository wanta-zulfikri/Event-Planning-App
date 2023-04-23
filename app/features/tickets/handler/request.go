package handler

type RequestCreateTicket struct {
	TicketType     string `json:"tickettype"`
	TicketCategory string `json:"ticketcategory"`
	TicketPrice    string `json:"ticketprice"`
	TicketQuantity string `json:"ticketquantity"`
}

type RequestUpdateTicket struct {
	TicketType     string `json:"tickettype"`
	TicketCategory string `json:"ticketcategory"`
	TicketPrice    string `json:"ticketprice"`
	TicketQuantity string `json:"ticketquantity"`
}
