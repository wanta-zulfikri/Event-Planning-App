package handler

type RequestWriteReview struct {
	Review string `from:"review"`
}

type RequestUpdateReview struct {
	Review string `from:"review"`
}
