package repository

import (
	"log"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances"
	"gorm.io/gorm"
)

type AttendancesRepository struct {
	db *gorm.DB 
} 

func New(db *gorm.DB) *AttendancesRepository {
	return &AttendancesRepository{db:db}
} 

func (ar *AttendancesRepository) CreateAttendance(newAttendance attendances.Core) (attendances.Core, error) {
	input := attendances.Attendances{
		ID: 			newAttendance.ID, 
		UserID: 		newAttendance.UserID, 
		EventID: 		newAttendance.EventID, 
		EventCategory:  newAttendance.EventCategory, 
		TicketType: 	newAttendance.TicketType, 
		Quantity:       newAttendance.Quantity,
	} 

	err := ar.db.Table("attendances").Create(&input).Error 
	if err != nil {
		log.Println("Error creating new attendances: ", err.Error()) 
		return attendances.Core{}, err 
	} 

	createdAttendances := attendances.Core{
		ID: 			input.ID, 
		UserID: 		input.UserID, 
		EventID: 		input.EventID, 
		EventCategory: 	input.EventCategory, 
		TicketType: 	input.TicketType, 
		Quantity: 		input.Quantity,
	} 
	return createdAttendances, nil 
} 

func (ar *AttendancesRepository) GetAttendance(id uint) (attendances.Core, error) {
	var input attendances.Attendances 
	result := ar.db.Where("id = ?", id).Find(&input) 
	if result.Error != nil {
		return attendances.Core{}, result.Error
	} 
	if result.RowsAffected == 0 {
		return attendances.Core{}, result.Error
	} 
	return attendances.Core{
		ID: 				input.ID, 
		UserID: 			input.UserID, 
		EventID: 			input.EventID,
		EventCategory: 		input.EventCategory, 
		TicketType: 		input.TicketType, 
		Quantity: 			input.Quantity,
	},nil

}