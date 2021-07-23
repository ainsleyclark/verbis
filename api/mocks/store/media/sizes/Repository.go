// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: mediaID, _a1
func (_m *Repository) Create(mediaID int, _a1 domain.MediaSizes) (domain.MediaSizes, error) {
	ret := _m.Called(mediaID, _a1)

	var r0 domain.MediaSizes
	if rf, ok := ret.Get(0).(func(int, domain.MediaSizes) domain.MediaSizes); ok {
		r0 = rf(mediaID, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.MediaSizes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.MediaSizes) error); ok {
		r1 = rf(mediaID, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: mediaID
func (_m *Repository) Delete(mediaID int) error {
	ret := _m.Called(mediaID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(mediaID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: mediaID
func (_m *Repository) Find(mediaID int) (domain.MediaSizes, error) {
	ret := _m.Called(mediaID)

	var r0 domain.MediaSizes
	if rf, ok := ret.Get(0).(func(int) domain.MediaSizes); ok {
		r0 = rf(mediaID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.MediaSizes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(mediaID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
