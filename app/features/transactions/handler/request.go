package handler

type RequestCreateTransaction struct {
	EventID         uint      `json:"event_id"`
	ItemDescription []Tickets `json:"item_description"`
	GrandTotal      uint      `json:"grandtotal"`
	PaymentMethod   string    `json:"payment_method"`
}

type Tickets struct {
	TicketID       uint   `json:"ticket_id"`
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
	Subtotal       uint   `json:"subtotal"`
}
