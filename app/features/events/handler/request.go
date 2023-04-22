package handler

type RequestCreateEvent struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	EventDate   string `form:"eventdate"`
	EventTime   string `form:"eventtime"`
	Status      string `form:"status"`
	Category    string `form:"category"`
	Location    string `form:"location"`
	Image       string `form:"image"`
}
