// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	builder "github.com/verbiscms/verbis/api/database/builder"

	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"

	version "github.com/hashicorp/go-version"
)

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

// Builder provides a mock function with given fields:
func (_m *Driver) Builder() *builder.Sqlbuilder {
	ret := _m.Called()

	var r0 *builder.Sqlbuilder
	if rf, ok := ret.Get(0).(func() *builder.Sqlbuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*builder.Sqlbuilder)
		}
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *Driver) DB() *sqlx.DB {
	ret := _m.Called()

	var r0 *sqlx.DB
	if rf, ok := ret.Get(0).(func() *sqlx.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.DB)
		}
	}

	return r0
}

// Drop provides a mock function with given fields:
func (_m *Driver) Drop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dump provides a mock function with given fields: path, filename
func (_m *Driver) Dump(path string, filename string) error {
	ret := _m.Called(path, filename)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(path, filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Install provides a mock function with given fields:
func (_m *Driver) Install() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrate provides a mock function with given fields: _a0
func (_m *Driver) Migrate(_a0 *version.Version) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*version.Version) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Schema provides a mock function with given fields:
func (_m *Driver) Schema() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Tables provides a mock function with given fields:
func (_m *Driver) Tables() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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
