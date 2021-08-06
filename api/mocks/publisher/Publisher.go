// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	domain "github.com/verbiscms/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"

	publisher "github.com/verbiscms/verbis/api/publisher"
)

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

// Asset provides a mock function with given fields: g, webp
func (_m *Publisher) Asset(g *gin.Context, webp bool) (*[]byte, domain.Mime, error) {
	ret := _m.Called(g, webp)

	var r0 *[]byte
	if rf, ok := ret.Get(0).(func(*gin.Context, bool) *[]byte); ok {
		r0 = rf(g, webp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]byte)
		}
	}

	var r1 domain.Mime
	if rf, ok := ret.Get(1).(func(*gin.Context, bool) domain.Mime); ok {
		r1 = rf(g, webp)
	} else {
		r1 = ret.Get(1).(domain.Mime)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gin.Context, bool) error); ok {
		r2 = rf(g, webp)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NotFound provides a mock function with given fields: g
func (_m *Publisher) NotFound(g *gin.Context) {
	_m.Called(g)
}

// Page provides a mock function with given fields: g
func (_m *Publisher) Page(g *gin.Context) ([]byte, error) {
	ret := _m.Called(g)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*gin.Context) []byte); ok {
		r0 = rf(g)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(g)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SiteMap provides a mock function with given fields:
func (_m *Publisher) SiteMap() publisher.SiteMapper {
	ret := _m.Called()

	var r0 publisher.SiteMapper
	if rf, ok := ret.Get(0).(func() publisher.SiteMapper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(publisher.SiteMapper)
		}
	}

	return r0
}

// Upload provides a mock function with given fields: g, webp
func (_m *Publisher) Upload(g *gin.Context, webp bool) (*[]byte, domain.Mime, error) {
	ret := _m.Called(g, webp)

	var r0 *[]byte
	if rf, ok := ret.Get(0).(func(*gin.Context, bool) *[]byte); ok {
		r0 = rf(g, webp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]byte)
		}
	}

	var r1 domain.Mime
	if rf, ok := ret.Get(1).(func(*gin.Context, bool) domain.Mime); ok {
		r1 = rf(g, webp)
	} else {
		r1 = ret.Get(1).(domain.Mime)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gin.Context, bool) error); ok {
		r2 = rf(g, webp)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
