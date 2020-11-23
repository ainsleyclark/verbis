// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// SiteRepository is an autogenerated mock type for the SiteRepository type
type SiteRepository struct {
	mock.Mock
}

// GetGlobalConfig provides a mock function with given fields:
func (_m *SiteRepository) GetGlobalConfig() *domain.Site {
	ret := _m.Called()

	var r0 *domain.Site
	if rf, ok := ret.Get(0).(func() *domain.Site); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Site)
		}
	}

	return r0
}

// GetLayouts provides a mock function with given fields:
func (_m *SiteRepository) GetLayouts() (domain.Layouts, error) {
	ret := _m.Called()

	var r0 domain.Layouts
	if rf, ok := ret.Get(0).(func() domain.Layouts); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.Layouts)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTemplates provides a mock function with given fields:
func (_m *SiteRepository) GetTemplates() (domain.Templates, error) {
	ret := _m.Called()

	var r0 domain.Templates
	if rf, ok := ret.Get(0).(func() domain.Templates); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.Templates)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThemeConfig provides a mock function with given fields:
func (_m *SiteRepository) GetThemeConfig() (domain.ThemeConfig, error) {
	ret := _m.Called()

	var r0 domain.ThemeConfig
	if rf, ok := ret.Get(0).(func() domain.ThemeConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.ThemeConfig)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
