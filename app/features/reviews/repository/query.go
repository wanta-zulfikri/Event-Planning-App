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
	input := reviews.Review{
		ID: 		 newReview.ID,
		UserID:      newReview.UserID,
		EventID:     newReview.EventID,   
		ReviewScore: newReview.ReviewScore,
		ReviewText:  newReview.ReviewText,
		
	}

	err := rr.db.Table("review").Create(&input).Error
	if err != nil {
		log.Println("Error creating new review: ", err.Error())
		return reviews.Core{}, err
	}

	createdReview := reviews.Core{
		ID: 		 input.ID,
		UserID:      input.UserID,
		EventID:     input.EventID,   
		ReviewScore: input.ReviewScore,
		ReviewText:  input.ReviewText, 
		
	}
	return createdReview, nil
} 

func (rr *ReviewRepository) UpdateReview(id uint, updateReview reviews.Core) error {
	input := reviews.Review{} 
	if err := rr.db.Where("id = ?", id).First(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		} 
		return err 

	} 

	input.ID = updateReview.ID 
	input.UserID = updateReview.UserID 
	input.EventID = updateReview.EventID 
	input.ReviewScore = updateReview.ReviewScore 
	input.ReviewText = updateReview.ReviewText 

	if err := rr.db.Save(&input).Error; err != nil {
		return err 
	} 
	return nil
}

func (rr *ReviewRepository) DeleteReview(id uint) error {
	input := reviews.Review{} 
	if err := rr.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		}
		return err 
	}  

	input.DeletedAt = gorm.DeletedAt{Time: time.Now(),Valid: true } 

	if err := rr.db.Save(&input).Error; err != nil {
			return err 
	} 
	return nil
}