// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: formID, f
func (_m *Repository) Create(formID int, f domain.FormSubmission) error {
	ret := _m.Called(formID, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, domain.FormSubmission) error); ok {
		r0 = rf(formID, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: formID
func (_m *Repository) Delete(formID int) error {
	ret := _m.Called(formID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(formID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: formID
func (_m *Repository) Find(formID int) (domain.FormSubmissions, error) {
	ret := _m.Called(formID)

	var r0 domain.FormSubmissions
	if rf, ok := ret.Get(0).(func(int) domain.FormSubmissions); ok {
		r0 = rf(formID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.FormSubmissions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(formID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
