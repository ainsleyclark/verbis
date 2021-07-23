// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	verbisfs "github.com/verbiscms/verbis/api/verbisfs"
)

// TemplateConfig is an autogenerated mock type for the TemplateConfig type
type TemplateConfig struct {
	mock.Mock
}

// GetExtension provides a mock function with given fields:
func (_m *TemplateConfig) GetExtension() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetFS provides a mock function with given fields:
func (_m *TemplateConfig) GetFS() verbisfs.FS {
	ret := _m.Called()

	var r0 verbisfs.FS
	if rf, ok := ret.Get(0).(func() verbisfs.FS); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(verbisfs.FS)
		}
	}

	return r0
}

// GetMaster provides a mock function with given fields:
func (_m *TemplateConfig) GetMaster() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetRoot provides a mock function with given fields:
func (_m *TemplateConfig) GetRoot() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
