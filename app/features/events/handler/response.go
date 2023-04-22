package handler 

import (

	"Event-Planning-App/app/features/events"
)
type AllEventResponse struct {
	ID 				 uint   `json:"id"` 
	Title			 string `json:"title"`
	Image			 string `json:"image"`
	Username	     string `json:"username
}

func CoreToGetAllEventRespB(data events.Core) AllEventResponse {
	return AllEventResponse{
		ID: 		data.Id, 
		Title: 		data.Title, 
		Image:		data.Image, 
		Username:   data.username,

	}
}

func CoreToGetAllEventResp(data []events.Core) []AllEventResponse {
	res := []AllEventResponse{}
	for _, val := range data {
		res = append(res,CoreToGetAllEventRespB(val))
	}
	return res
} 

type MyEventResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"Image"`
} 

func CoreToGetMyEvent(data events.Core) MyEventResponse {
	return MyEventResponse{
		ID: 		data.ID, 
		Title: 		data.Title, 
		Image:		data.Image, 
	}
} 

func CoreToGetAllEventResp(data []events.Core) []AllEventResponse {
		res:= []MyEventResponse{} 
		for _, val := range data{
			res = append(res, CoreToGetMyEvent(val))
		} 
		return res	
}	