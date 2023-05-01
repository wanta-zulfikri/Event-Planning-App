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
	input := Review{
		UserID:    request.UserID,
		Username:  request.Username,
		EventID:   request.EventID,
		Review:    request.Review,
		UpdatedAt: time.Now(),
	}

	input.Username = request.Username
	input.EventID = request.EventID
	input.Review = request.Review

	if err := rr.db.Save(&input).Error; err != nil {
		return reviews.Core{}, err
	}

	if err := rr.db.Model(&Review{}).Where("id = ? AND deleted_at IS NULL", request.EventID).Updates(Review{Review: input.Review, UpdatedAt: time.Now()}).Error; err != nil {
		log.Println("Error updating review: ", err.Error())
		return reviews.Core{}, err
	}

	updateReview := reviews.Core{
		UserID:   request.UserID,
		Username: request.Username,
		EventID:  request.EventID,
		Review:   input.Review,
	}

	return updateReview, nil
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
