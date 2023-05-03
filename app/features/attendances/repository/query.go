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
	return &AttendancesRepository{db: db}
}

func (ar *AttendancesRepository) CreateAttendance(newAttendance attendances.Core) (attendances.Core, error) {
	input := Attendance{
		EventID:       newAttendance.EventID,
		Title:         newAttendance.Title,
		Description:   newAttendance.Description,
		HostedBy:      newAttendance.HostedBy,
		Date:          newAttendance.Date,
		Time:          newAttendance.Time,
		Status:        newAttendance.Status,
		Category:      newAttendance.Category,
		Location:      newAttendance.Location,
		EventPicture:  newAttendance.EventPicture,
	}

	err := ar.db.Table("attendances").Create(&input).Error
	if err != nil {
		log.Println("Error creating new attendances: ", err.Error())
		return attendances.Core{}, err
	}

	createdAttendances := attendances.Core{ 
		Description:   input.Description,
		HostedBy:      input.HostedBy,
		Date:          input.Date,
		Time:          input.Time,
		Status:        input.Status,
		Category:      input.Category, 
		Location:      input.Location, 
		EventPicture:  input.EventPicture,

	}
	return createdAttendances, nil
}

func (ar *AttendancesRepository) GetAttendance() ([]attendances.Core, error){
	var cores []attendances.Core
	if err := ar.db.Table("reviews").Where("deleted_at IS NULL").Find(&cores).Error; err != nil {
	  return nil, err 
	} 
	return cores, nil 
}