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

// Exists provides a mock function with given fields: _a0
func (_m *Repository) Exists(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Find provides a mock function with given fields: _a0
func (_m *Repository) Find(_a0 string) (*domain.ThemeConfig, error) {
	ret := _m.Called(_a0)

	var r0 *domain.ThemeConfig
	if rf, ok := ret.Get(0).(func(string) *domain.ThemeConfig); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ThemeConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Layouts provides a mock function with given fields: _a0
func (_m *Repository) Layouts(_a0 string) (domain.Layouts, error) {
	ret := _m.Called(_a0)

	var r0 domain.Layouts
	if rf, ok := ret.Get(0).(func(string) domain.Layouts); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Layouts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: activeTheme
func (_m *Repository) List(activeTheme string) ([]*domain.ThemeConfig, error) {
	ret := _m.Called(activeTheme)

	var r0 []*domain.ThemeConfig
	if rf, ok := ret.Get(0).(func(string) []*domain.ThemeConfig); ok {
		r0 = rf(activeTheme)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.ThemeConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(activeTheme)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Screenshot provides a mock function with given fields: _a0, file
func (_m *Repository) Screenshot(_a0 string, file string) ([]byte, domain.Mime, error) {
	ret := _m.Called(_a0, file)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(_a0, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 domain.Mime
	if rf, ok := ret.Get(1).(func(string, string) domain.Mime); ok {
		r1 = rf(_a0, file)
	} else {
		r1 = ret.Get(1).(domain.Mime)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(_a0, file)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Templates provides a mock function with given fields: _a0
func (_m *Repository) Templates(_a0 string) (domain.Templates, error) {
	ret := _m.Called(_a0)

	var r0 domain.Templates
	if rf, ok := ret.Get(0).(func(string) domain.Templates); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Templates)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
