package handler 


type RequestWriteReview struct {
	ID      	uint `from:"id"`   
	UserID  	uint `from:"user_id"`
	EventID 	uint `from:"event_id"`  
	ReviewScore  int `from:"review_score"`
	ReviewText string `from:"review_text"`

}


type RequestUpdateReview struct {
	ID      	uint `from:"id"`   
	UserID  	uint `from:"user_id"`
	EventID 	uint `from:"event_id"`  
	ReviewScore  int `from:"review_score"`
	ReviewText string `from:"review_text"`

}