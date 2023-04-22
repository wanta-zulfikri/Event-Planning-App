package repository

import (
	"Event-Planning-App/app/features/events"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model 
	UserID 			uint 
	Title 			string 
	Details 		string
	Hosted_by      	string 
	Date            string  
	Time            string  
	Status 			string 
	Event_category  string 
	Location 		string 
	Image 			string    	
} 


func CoreToEvent(data events.Core) Event {
	return Event{
		Model: 				gorm.Model{ID: data.Id}, 
		Title: 				data.Title, 
		Details: 			data.Details, 
		Hosted_by: 			data.Hosted_by, 
		Date: 				data.Date,
		Time:				data.Time, 
		Status: 			data.Status, 
		Event_category: 	data.Event_category,
		Location: 			data.Location,	
		Image: 				data.Image,			
	}
}

func EventToCore(data Event) events.Core {
	return events.Core{
		Id: 				data.ID, 
		Title:          	data.Title, 
		Details: 			data.Details, 
		Hosted_by:      	data.Hosted_by, 
		Date:				data.Date, 
		Time:   			data.Time, 
		Status:				data.Status, 
		Event_category: 	data.Event_category, 
		Location:       	data.Location, 
		Image:				data.Image, 
	}
} 

func ListModelToCore(dataModel []Event) []events.Core {
	var dataCore []events.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, EventToCore(v))
	}
	return dataCore
}