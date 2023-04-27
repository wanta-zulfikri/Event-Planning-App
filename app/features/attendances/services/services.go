package services

import "github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances"

type AttendancesService struct {
	a attendances.Repository
}  

func New(z attendances.Repository) attendances.Service {
	return &AttendancesService{a:z}
} 

func (as *AttendancesService) CreateAttendance(newAttendance attendances.Core) error {
	_, err := as.a.CreateAttendance(newAttendance) 
	if err != nil {
		return err 
	} 
	return nil 
} 

func (as *AttendancesService)GetAttendance() ([]attendances.Core, error) {
	attendance, err := as.a.GetAttendance() 
	if err != nil { 
		return nil, err 
	} 
	return attendance, nil		
}