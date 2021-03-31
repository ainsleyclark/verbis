// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	bytes "bytes"

	mock "github.com/stretchr/testify/mock"
)

// Minifier is an autogenerated mock type for the Minifier type
type Minifier struct {
	mock.Mock
}

// Minify provides a mock function with given fields: name, mime
func (_m *Minifier) Minify(name string, mime string) ([]byte, error) {
	ret := _m.Called(name, mime)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(name, mime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, mime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinifyBytes provides a mock function with given fields: b, mime
func (_m *Minifier) MinifyBytes(b *bytes.Buffer, mime string) ([]byte, error) {
	ret := _m.Called(b, mime)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*bytes.Buffer, string) []byte); ok {
		r0 = rf(b, mime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bytes.Buffer, string) error); ok {
		r1 = rf(b, mime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
