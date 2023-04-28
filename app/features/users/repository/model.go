package repository

import (
	events "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/repository"
	reviews "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews/repository"
	transacations "github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions/repository"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string                      `json:"username" gorm:"type:varchar(100);not null"`
	Email        string                      `json:"email" gorm:"type:varchar(100);not null"`
	Phone        string                      `json:"phone" gorm:"type:varchar(15);not null"`
	Password     string                      `json:"password" gorm:"type:varchar(100);not null"`
	Image        string                      `json:"image" gorm:"type:varchar(100)"`
	Events       []events.Event              `gorm:"foreignKey:UserID"`
	Transactions []transacations.Transaction `gorm:"foreignKey:UserID"`
	Reviews      []reviews.Review            `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Image == "" {
		u.Image = "https://storage.googleapis.com/grouproject/book/images.jpeg"
	}
	return nil
}
