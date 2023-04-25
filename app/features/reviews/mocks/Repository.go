// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	reviews "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteReview provides a mock function with given fields: id
func (_m *Repository) DeleteReview(id uint) error {
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
func (_m *Repository) UpdateReview(id uint, updateReview reviews.Core) error {
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
func (_m *Repository) WriteReview(newReview reviews.Core) (reviews.Core, error) {
	ret := _m.Called(newReview)

	var r0 reviews.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(reviews.Core) (reviews.Core, error)); ok {
		return rf(newReview)
	}
	if rf, ok := ret.Get(0).(func(reviews.Core) reviews.Core); ok {
		r0 = rf(newReview)
	} else {
		r0 = ret.Get(0).(reviews.Core)
	}

	if rf, ok := ret.Get(1).(func(reviews.Core) error); ok {
		r1 = rf(newReview)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
