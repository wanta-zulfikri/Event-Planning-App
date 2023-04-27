package handler 

type RequestCreateAttendances struct {
	UserID        uint 		`from:"user_id"`
	EventID       uint 		`from:"event_id"`
	EventCategory string 	`from:"event_category"`
	TicketType    string	`from:"ticket_type"`
	Quantity      string    `from:"quantity"`
} 

type RequestGetAttendances struct {
	ID            uint 		`from:"id"`
	UserID        uint 		`from:"user_id"`
	EventID       uint 		`from:"event_id"`
	EventCategory string 	`from:"event_category"`
	TicketType    string	`from:"ticket_type"`
	Quantity      string    `from:"quantity"`
} 
