package handler 

type RequestCreateAttendances struct { 
	ID             uint         `json:"user_id"`
	EventID        uint 		`json:"event_id"`
	Title          string 	    `json:"title"`
	Description    string	    `json:"description"`
	HostedBy       string       `json:"hosted_by"`
	Date           string       `json:"date"` 
	Time           string       `json:"time"`
	Status         string       `json:"status"`
	Location       string       `json:"location"`
	EventPicture   string       `json:"event_picture"` 
	Category       string       `json:"category"`
} 

type RequestGetAttendances struct {
	ID             uint         `json:"user_id"`
	EventID        uint 		`json:"event_id"`
	Title          string 	    `json:"title"`
	Description    string	    `json:"description"`
	HostedBy       string       `json:"hosted_by"`
	Date           string       `json:"date"` 
	Time           string       `json:"time"`
	Status         string       `json:"status"`
	Location       string       `json:"location"`
	EventPicture   string       `json:"event_picture"`
	Category       string       `json:"category"`
} 
