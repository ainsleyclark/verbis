// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	render "github.com/ainsleyclark/verbis/api/render"
)

// Renderer is an autogenerated mock type for the Renderer type
type Renderer struct {
	mock.Mock
}

// Asset provides a mock function with given fields: g
func (_m *Renderer) Asset(g *gin.Context) (string, *[]byte, error) {
	ret := _m.Called(g)

	var r0 string
	if rf, ok := ret.Get(0).(func(*gin.Context) string); ok {
		r0 = rf(g)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *[]byte
	if rf, ok := ret.Get(1).(func(*gin.Context) *[]byte); ok {
		r1 = rf(g)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*[]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gin.Context) error); ok {
		r2 = rf(g)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NotFound provides a mock function with given fields: g
func (_m *Renderer) NotFound(g *gin.Context) {
	_m.Called(g)
}

// Page provides a mock function with given fields: g
func (_m *Renderer) Page(g *gin.Context) ([]byte, error) {
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
func (_m *Renderer) SiteMap() render.SiteMapper {
	ret := _m.Called()

	var r0 render.SiteMapper
	if rf, ok := ret.Get(0).(func() render.SiteMapper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(render.SiteMapper)
		}
	}

	return r0
}

// Upload provides a mock function with given fields: g
func (_m *Renderer) Upload(g *gin.Context) (string, *[]byte, error) {
	ret := _m.Called(g)

	var r0 string
	if rf, ok := ret.Get(0).(func(*gin.Context) string); ok {
		r0 = rf(g)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *[]byte
	if rf, ok := ret.Get(1).(func(*gin.Context) *[]byte); ok {
		r1 = rf(g)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*[]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gin.Context) error); ok {
		r2 = rf(g)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
