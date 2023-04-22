package handler

type EventRequest struct {
	UserID 			uint    `json:"user_id" from:"user_id` 
	Title 			string  `json:"title" from:"title"`
	Details 		string  `json:"details" from:"details"`
	Hosted_by      	string  `json:"hosted_by" from:"hosted_by"`
	Date            string  `json:"date" from:"date"`
	Time            string  `json:"time" from:"time"`
	Status 			string  `json:"status" from:"status"`
	Event_category  string  `json:"event_category" from:"event_category"`
	Location 		string  `json:"location" from:"location"`
	Image 			string 	`json:"image" from:"image"`
}