package handler

type RequestCreateTransaction struct {
	ItemDescription []Tickets `json:"item_description"`
}

type Tickets struct {
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
}
