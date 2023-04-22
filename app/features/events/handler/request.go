package handler

import "time"

type CreateEvent struct {
	Title       string    `form:"title"`
	Description string    `form:"description"`
	EventDate   time.Time `form:"eventdate"`
	EventTime   time.Time `form:"eventtime"`
	Status      string    `form:"status"`
	Category    string    `form:"category"`
	Location    string    `form:"location"`
	Image       string    `form:"image"`
}
