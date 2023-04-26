package handler 


type ResponseWriteReview struct {
	Review   string  `from:"review"` 
	Username string  `from:"username"` 
	Image    string  `from:"image"`
}

type ResponseUpdateReview struct {
	Review   string  `from:"review"` 
	Username string  `from:"username"`  
	Image 	 string  `from:"image"`
}
