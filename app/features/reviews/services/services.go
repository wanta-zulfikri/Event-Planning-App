package services

import (
	"errors"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews"
	"gorm.io/gorm"
)

type ReviewService struct {
	n reviews.Repository
}

func New(o reviews.Repository) reviews.Service {
	return &ReviewService{n: o}
}

func (rs *ReviewService) WriteReview(newReview reviews.Core) error {
	_, err := rs.n.WriteReview(newReview)
	if err != nil {
		return err
	}
	return nil
} 

func (rs *ReviewService)UpdateReview(id uint, updateReview reviews.Core) error {
	updateReview.ID = id 
	if err := rs.n.UpdateReview(id, updateReview); err != nil {
		return err
	} 
	return nil 
} 

func (rs *ReviewService) DeleteReview(id uint) error {
	err := rs.n.DeleteReview(id) 
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		} 
		return err 
	} 
	return nil
}
