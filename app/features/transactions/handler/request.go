package handler

type RequestCreateTransaction struct {
	ItemDescription []Tickets `json:"item_description"`
}

type Tickets struct {
	TicketID uint `json:"ticket_id"`
}
