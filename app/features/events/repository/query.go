package repository

import (
	"Event-Planning-App/app/features/events"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type eventQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) events.EventData {
	return &eventQuery{
		db: db,
	}
}


func (ev *eventQuery) MyEvent(userid int, limit int, offset int) ([]events.Core, error) {
	var eventsModel []Event 
	tx := ev.db.Limit(limit).Offset(offset).Where("user_id = ?", userid).Find(&eventsModel) 
	if tx.Error != nil { 
		log.Error("Terjadi error saat select event") 
		return nil, tx.Error
	}
	eventsCoreAll := ListModelToCore(eventsModel) 
	return eventsCoreAll, nil
} 

func (ev *eventQuery) SelectAll(limit int, offset int, name string) ([]events.Core, error) {
	nameSearch := "%" + name + "%" 
	var eventsModel []Event 
	tx := ev.db.Limit(limit).Offset(offset).Where("events.title LIKE ?", nameSearch).Select("events.id, events.title,events.Image, users.name AS user_name").Joins("JOIN users ON events.user_id = users.id").Group("events.id").Find(&eventsModel) 
	if tx.Error != nil { 
		log.Error("Terjadi error saat select event") 
		return nil, tx.Error
	} 
	eventsCoreAll := ListModelToCore(eventsModel) 
	return eventsCoreAll, nil
} 

func (ev *eventQuery) Insert(input events.Core) error {
	data := CoreToEvent(input) 
	tx := ev.db.Create(&data) 
	if tx.Error != nil {
		log.Error("Terjadi error saat Insert event") 
		return tx.Error
	} 
	return nil 
}  


func (ev *eventQuery) Update(userId uint, id uint, input events.Core) error {
	data := CoreToEvent(input)  
	tx := ev.db.Model(&Event{}).Where("id = ? AND user_id = ?", id, userId).Updates(&data)  
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error saat update event") 
		return errors.New("event no updated")
	} 
	if tx.Error != nil {
		log.Error("event tidak ditemukan") 
		return tx.Error
	}
	return nil 
} 

func (ev *eventQuery) GetEventById(id uint) (events.Core, error) {
	tmp := Event{} 
	tx := ev.db.Where("id = ?", id).First(&tmp) 
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error saat select event ") 
		return events.Core{}, errors.New("event not found")
	} 
	if tx.Error != nil {
		log.Error("event tidak ditemukan") 
		return events.Core{}, tx.Error
	}

	return EventToCore(tmp), nil 
} 

func (ev *eventQuery) DeleteEvent(userid int, id int) error {
	tx := ev.db.Where("user_id = ?", userid).Delete(&Event{}, id) 
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error")
		return errors.New("no event deleted")
	} 
		if tx.Error != nil {
		log.Error("event tidak ditemukan") 
		return tx.Error
	}
	return nil 
}

