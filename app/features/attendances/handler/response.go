package handler 


type attendancesResponse struct {
	Code       int                      `json:"code"`
	Message    string                   `json:"message"`
	Data       []RequestGetAttendances  `json:"data"`
	Pagination Pagination               `json:"pagination"`
}

type Pagination struct {
	Page       int                      `json:"page"`
	PerPage    int                      `json:"per_page"`
	TotalPages int                      `json:"total_pages"`
	TotalItems int                      `json:"total_items"`
}   



type ResponseGetAttendances struct {
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