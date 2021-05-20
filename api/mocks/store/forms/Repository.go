// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"

	params "github.com/ainsleyclark/verbis/api/helpers/params"

	uuid "github.com/google/uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: f
func (_m *Repository) Create(f domain.Form) (domain.Form, error) {
	ret := _m.Called(f)

	var r0 domain.Form
	if rf, ok := ret.Get(0).(func(domain.Form) domain.Form); ok {
		r0 = rf(f)
	} else {
		r0 = ret.Get(0).(domain.Form)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Form) error); ok {
		r1 = rf(f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: id
func (_m *Repository) Find(id int) (domain.Form, error) {
	ret := _m.Called(id)

	var r0 domain.Form
	if rf, ok := ret.Get(0).(func(int) domain.Form); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Form)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUUID provides a mock function with given fields: uniq
func (_m *Repository) FindByUUID(uniq uuid.UUID) (domain.Form, error) {
	ret := _m.Called(uniq)

	var r0 domain.Form
	if rf, ok := ret.Get(0).(func(uuid.UUID) domain.Form); ok {
		r0 = rf(uniq)
	} else {
		r0 = ret.Get(0).(domain.Form)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(uniq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: meta
func (_m *Repository) List(meta params.Params) (domain.Forms, int, error) {
	ret := _m.Called(meta)

	var r0 domain.Forms
	if rf, ok := ret.Get(0).(func(params.Params) domain.Forms); ok {
		r0 = rf(meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Forms)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(params.Params) int); ok {
		r1 = rf(meta)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(params.Params) error); ok {
		r2 = rf(meta)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Submit provides a mock function with given fields: f
func (_m *Repository) Submit(f domain.FormSubmission) error {
	ret := _m.Called(f)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.FormSubmission) error); ok {
		r0 = rf(f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: f
func (_m *Repository) Update(f domain.Form) (domain.Form, error) {
	ret := _m.Called(f)

	var r0 domain.Form
	if rf, ok := ret.Get(0).(func(domain.Form) domain.Form); ok {
		r0 = rf(f)
	} else {
		r0 = ret.Get(0).(domain.Form)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Form) error); ok {
		r1 = rf(f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
