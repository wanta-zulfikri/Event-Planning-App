// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	reviews "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// DeleteReview provides a mock function with given fields: id
func (_m *Service) DeleteReview(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateReview provides a mock function with given fields: id, updateReview
func (_m *Service) UpdateReview(id uint, updateReview reviews.Core) error {
	ret := _m.Called(id, updateReview)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, reviews.Core) error); ok {
		r0 = rf(id, updateReview)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteReview provides a mock function with given fields: newReview
func (_m *Service) WriteReview(newReview reviews.Core) error {
	ret := _m.Called(newReview)

	var r0 error
	if rf, ok := ret.Get(0).(func(reviews.Core) error); ok {
		r0 = rf(newReview)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
