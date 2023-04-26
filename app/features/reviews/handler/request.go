package handler

type RequestWriteReview struct {
	EventID  uint    `from:"event_id"`
	Review   string  `from:"review"` 
}

type RequestUpdateReview struct {
	
	Review   string `from:"review"` 
	
}
