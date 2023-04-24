package handler

type ResponseGetTickets struct {
	TicketType     string `json:"tickettype"`
	TicketCategory string `json:"ticketcategory"`
	TicketPrice    uint   `json:"ticketprice"`
	TicketQuantity uint   `json:"ticketquantity"`
}
