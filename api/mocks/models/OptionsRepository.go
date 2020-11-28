// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// OptionsRepository is an autogenerated mock type for the OptionsRepository type
type OptionsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, value
func (_m *OptionsRepository) Create(name string, value interface{}) error {
	ret := _m.Called(name, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(name, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: name
func (_m *OptionsRepository) Exists(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *OptionsRepository) Get() (domain.OptionsDB, error) {
	ret := _m.Called()

	var r0 domain.OptionsDB
	if rf, ok := ret.Get(0).(func() domain.OptionsDB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.OptionsDB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *OptionsRepository) GetByName(name string) (interface{}, error) {
	ret := _m.Called(name)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStruct provides a mock function with given fields:
func (_m *OptionsRepository) GetStruct() domain.Options {
	ret := _m.Called()

	var r0 domain.Options
	if rf, ok := ret.Get(0).(func() domain.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.Options)
	}

	return r0
}

// Update provides a mock function with given fields: name, value
func (_m *OptionsRepository) Update(name string, value interface{}) error {
	ret := _m.Called(name, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(name, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCreate provides a mock function with given fields: options
func (_m *OptionsRepository) UpdateCreate(options *domain.OptionsDB) error {
	ret := _m.Called(options)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.OptionsDB) error); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
