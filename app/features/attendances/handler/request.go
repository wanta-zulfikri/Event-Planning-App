package handler 

type RequestCreateAttendances struct {
	EventID        uint 		`json:"event_id"`
	Title          string 	    `json:"title"`
	Description    string	    `json:"description"`
	HostedBy       string       `json:"hosted_by"`
	Date           string       `json:"date"` 
	Time           string       `json:"time"`
	Status         string       `json:"status"`
	Location       string       `json:"location"`
	EventPicture   string       `json:"event_picture"`
} 

type RequestGetAttendances struct {
	EventID        uint 		`json:"event_id"`
	Title          string 	    `json:"title"`
	Description    string	    `json:"description"`
	HostedBy       string       `json:"hosted_by"`
	Date           string       `json:"date"` 
	Time           string       `json:"time"`
	Status         string       `json:"status"`
	Location       string       `json:"location"`
	EventPicture   string       `json:"event_picture"`
} 
