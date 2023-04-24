package handler

type ResponseGetEvents struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"eventdate"`
	EventTime   string `json:"eventtime"`
	Status      string `json:"status"`
	Category    string `json:"category"`
	Location    string `json:"location"`
	Image       string `json:"image"`
	Hostedby    uint   `json:"hostedby"`
}
