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

func (rs *ReviewService) WriteReview(request reviews.Core) (reviews.Core, error) {
	result, err := rs.n.WriteReview(request)
	if err != nil {
		return request, errors.New(err.Error())
	}
	return result, nil
}

func (rs *ReviewService) UpdateReview(request reviews.Core) (reviews.Core, error) {
	result, err := rs.n.UpdateReview(request)
	if err != nil {
		return request, errors.New(err.Error())
	}
	return result, nil
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
