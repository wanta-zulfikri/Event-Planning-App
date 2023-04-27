package handler

type ResponseGetEvents struct {
	ID            uint   `json:"event_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Hosted_by     string `json:"hosted_by"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Status        string `json:"status"`
	Category      string `json:"category"`
	Location      string `json:"location"`
	Event_picture string `json:"event_picture"`
}

type ResponseGetEvent struct {
	ID            uint   `json:"event_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Hosted_by     string `json:"hosted_by"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Status        string `json:"status"`
	Category      string `json:"category"`
	Location      string `json:"location"`
	Event_picture string `json:"event_picture"`
}

type EventResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    EventData `json:"data"`
}

type EventData struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	HostedBy    string           `json:"hosted_by"`
	Date        string           `json:"date"`
	Time        string           `json:"time"`
	Status      string           `json:"status"`
	Category    string           `json:"category"`
	Location    string           `json:"location"`
	Picture     string           `json:"event_picture"`
	Ticket      []TicketResponse `json:"ticket"`
}

type TicketResponse struct {
	Category string `json:"ticket_category"`
	Price    uint   `json:"ticket_price"`
	Quantity uint   `json:"ticket_quantity"`
}

type EventsResponse struct {
	Code       int                 `json:"code"`
	Message    string              `json:"message"`
	Data       []ResponseGetEvents `json:"data"`
	Pagination Pagination          `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type ResponseUpdateEvent struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Hosted_by     string `json:"hosted_by"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Status        string `json:"status"`
	Category      string `json:"category"`
	Location      string `json:"location"`
	Event_picture string `json:"event_picture"`
}
