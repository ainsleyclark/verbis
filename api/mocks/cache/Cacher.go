// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Cacher is an autogenerated mock type for the Cacher type
type Cacher struct {
	mock.Mock
}

// Flush provides a mock function with given fields:
func (_m *Cacher) Flush() {
	_m.Called()
}

// Get provides a mock function with given fields: k
func (_m *Cacher) Get(k string) (interface{}, bool) {
	ret := _m.Called(k)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(k)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(k)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Set provides a mock function with given fields: k, x, d
func (_m *Cacher) Set(k string, x interface{}, d time.Duration) {
	_m.Called(k, x, d)
}
