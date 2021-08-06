// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/verbiscms/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"
)

// Finder is an autogenerated mock type for the Finder type
type Finder struct {
	mock.Mock
}

// Layout provides a mock function with given fields: themePath, post, cacheable
func (_m *Finder) Layout(themePath string, post domain.PostDatum, cacheable bool) domain.FieldGroups {
	ret := _m.Called(themePath, post, cacheable)

	var r0 domain.FieldGroups
	if rf, ok := ret.Get(0).(func(string, domain.PostDatum, bool) domain.FieldGroups); ok {
		r0 = rf(themePath, post, cacheable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.FieldGroups)
		}
	}

	return r0
}
