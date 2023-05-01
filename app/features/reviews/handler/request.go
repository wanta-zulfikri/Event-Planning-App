package handler

type RequestWriteReview struct {
<<<<<<< HEAD
	EventID  uint    `from:"event_id"`
	Review   string  `from:"review"` 
}

type RequestUpdateReview struct {
	Review   string `from:"review"` 
=======
	Review string `from:"review"`
}

type RequestUpdateReview struct {
	Review string `from:"review"`
>>>>>>> feature/revisedreviews
}
