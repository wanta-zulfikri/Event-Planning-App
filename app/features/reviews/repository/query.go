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

func (rr *ReviewRepository) WriteReview(request reviews.Core) (reviews.Core, error) {
	input := Review{
		UserID:   request.UserID,
		Username: request.Username,
		EventID:  request.EventID,
		Review:   request.Review,
	}

	err := rr.db.Table("reviews").Create(&input).Error
	if err != nil {
		log.Println("Error creating new review: ", err.Error())
		return reviews.Core{}, err
	}

	createdReview := reviews.Core{
		Username: request.Username,
		EventID:  input.EventID,
		Review:   input.Review,
	}
	return createdReview, nil
}

func (rr *ReviewRepository) UpdateReview(request reviews.Core) (reviews.Core, error) {
	input := Review{}
	if err := rr.db.Where("id = ?", request.UserID).First(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return request, errors.New(err.Error())
		}
		return request, nil

	}

	input.Username = request.Username
	input.EventID = request.EventID
	input.Review = request.Review

	if err := rr.db.Save(&input).Error; err != nil {
		return request, errors.New(err.Error())
	}
	return request, nil
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
