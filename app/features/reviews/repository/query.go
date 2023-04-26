package repository

import (
	"errors"
	"log"
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (rr *ReviewRepository) WriteReview(newReview reviews.Core) (reviews.Core, error) {
	input := Review{
		UserID:   newReview.UserID,
		EventID:  newReview.EventID,
		Review:   newReview.Review, 
		Username: newReview.Username, 
		Image:    newReview.Image,
	}

	err := rr.db.Table("review").Create(&input).Error
	if err != nil {
		log.Println("Error creating new review: ", err.Error())
		return reviews.Core{}, err
	}

	createdReview := reviews.Core{
		UserID:  	input.UserID,
		EventID: 	input.EventID,
		Review:  	input.Review, 
		Username: 	input.Username,
		Image: 		input.Image,
	}
	return createdReview, nil
}

func (rr *ReviewRepository) UpdateReview(id uint, updateReview reviews.Core) error {
	input := Review{}
	if err := rr.db.Where("id = ?", id).First(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err

	}

	input.UserID = updateReview.UserID
	input.EventID = updateReview.EventID
	input.Review = updateReview.Review 
	input.Username = updateReview.Username 
	input.Image = updateReview.Image

	if err := rr.db.Save(&input).Error; err != nil {
		return err
	}
	return nil
}

func (rr *ReviewRepository) DeleteReview(id uint) error {
	input := Review{}
	if err := rr.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := rr.db.Save(&input).Error; err != nil {
		return err
	}
	return nil
}
