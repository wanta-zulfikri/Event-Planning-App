package handler

type ResponseGetTickets struct {
	TicketType     string `json:"tickettype"`
	TicketCategory string `json:"ticketcategory"`
	TicketPrice    string `json:"ticketprice"`
	TicketQuantity string `json:"ticketquantity"`
}
