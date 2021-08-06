// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/verbiscms/verbis/api/domain"
)

// System is an autogenerated mock type for the System type
type System struct {
	mock.Mock
}

// HasUpdate provides a mock function with given fields:
func (_m *System) HasUpdate() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Install provides a mock function with given fields: db, restart
func (_m *System) Install(db domain.InstallVerbis, restart bool) error {
	ret := _m.Called(db, restart)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.InstallVerbis, bool) error); ok {
		r0 = rf(db, restart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LatestVersion provides a mock function with given fields:
func (_m *System) LatestVersion() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Restart provides a mock function with given fields:
func (_m *System) Restart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: restart
func (_m *System) Update(restart bool) (string, error) {
	ret := _m.Called(restart)

	var r0 string
	if rf, ok := ret.Get(0).(func(bool) string); ok {
		r0 = rf(restart)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(restart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateInstall provides a mock function with given fields: step, install
func (_m *System) ValidateInstall(step int, install domain.InstallVerbis) error {
	ret := _m.Called(step, install)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, domain.InstallVerbis) error); ok {
		r0 = rf(step, install)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
