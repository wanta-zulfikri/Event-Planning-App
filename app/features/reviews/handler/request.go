package handler

type RequestWriteReview struct {
	ID      uint   `from:"id"`
	UserID  uint   `from:"user_id"`
	EventID uint   `from:"event_id"`
	Review  string `from:"review"`
}

type RequestUpdateReview struct {
	ID      uint   `from:"id"`
	UserID  uint   `from:"user_id"`
	EventID uint   `from:"event_id"`
	Review  string `from:"review"`
}
