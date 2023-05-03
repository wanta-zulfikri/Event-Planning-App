package handler

type ResponseGetTickets struct {
	EventID        uint   `json:"event_id"`
	TicketID       uint   `json:"ticket_id"`
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
}

type TicketResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    []ResponseGetTickets `json:"data"`
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}
