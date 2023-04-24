package repository

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events/repository"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string             `json:"username" gorm:"type:varchar(100);not null"`
	Email    string             `json:"email" gorm:"primaryKey"`
	Password string             `json:"password" gorm:"type:varchar(100);not null"`
	Image    string             `json:"image" gorm:"type:varchar(100)"`
	Events   []repository.Event `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Image == "" {
		u.Image = "https://storage.googleapis.com/grouproject/book/images.jpeg"
	}
	return nil
}
