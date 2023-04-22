package services

import (
	"Event-Planning-App/app/features/events"
	"Event-Planning-App/helper"
	"mime/multipart"

	"github.com/go-playground/validator/v10"
)

type eventService struct {
	data events.EventData
	vld  *validator.Validate
}


func New(repo events.EventService) events.EventService {
	return &eventService{
		   data: repo,
		   vld:  validator.New(),
	}
}

func (ev *eventService) Add(newEvent events.Core, fileHeader *multipart.FileHeader) error {
	errValidate := ev.vld.Struct(newEvent)

	if errValidate != nil {
		return errValidate
	}

	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.StorageGCP(file,)
		if err != nil {
			return err
		}
		newEvent.Image = uploadURL[0]
	}
	errInsert := ev.data.Insert(newEvent)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (ev *eventService) MyEvent(userid int, page int) ([]events.Core, error) {
	limit := 10
	offset := (page - 1) * limit
	data, err := ev.data.MyEvent(userid, limit, offset)
	return data, err
}

func (ev *eventService) Update(userid int, id int, updateEvent events.Core, fileHeader *multipart.FileHeader) error {
	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.StorageGCP(file, "/events")
		if err != nil {
			return err
		}
		updateEvent.Image = uploadURL[0]
	}

	errUpdate := ev.data.Update(uint(userid), uint(id), updateEvent)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
} 

func (ev *eventService) GetAll(page int, name string) ([]events.Core, error) {
	limit := 10 
	offset := (page - 1) * limit 
	data, err := ev.data.SelectAll(limit, offset, name)
	return data, err
}

func (ev *eventService) GetEventById(id int) (events.Core, error) {
	data, err := ev.data.GetEventById(uint(id))
	return data, err
}

func (ev *eventService) DeleteEvent(userid int, id int) error {
	errDelete := ev.data.DeleteEvent(userid, id) 
	if errDelete != nil {
		return errDelete
	}
	return nil
}