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