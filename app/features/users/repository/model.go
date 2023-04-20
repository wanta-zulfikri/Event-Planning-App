package repository

import (
	"Event-Planning-App/app/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"primaryKey;type:varchar(100);unique"`
	Password string `gorm:"type:varchar(100);not null"`
	Image    string `gorm:"type:varchar(30);not null"`
}

func CoreToUser(repository users.Core) User {
	return User{
		Model:    gorm.Model{ID: repository.ID},
		Username: repository.Username,
		Email:    repository.Email,
		Password: repository.Password,
	}
}

func UserToCore(repository User) users.Core {
	return users.Core{
		ID:       repository.ID,
		Username: repository.Username,
		Email:    repository.Email,
		Password: repository.Password,
	}
}
