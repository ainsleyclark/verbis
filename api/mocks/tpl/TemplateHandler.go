// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	domain "github.com/verbiscms/verbis/api/domain"

	io "io"

	mock "github.com/stretchr/testify/mock"

	template "html/template"

	tpl "github.com/verbiscms/verbis/api/tpl"
)

// TemplateHandler is an autogenerated mock type for the TemplateHandler type
type TemplateHandler struct {
	mock.Mock
}

// Data provides a mock function with given fields: ctx, post
func (_m *TemplateHandler) Data(ctx *gin.Context, post *domain.PostDatum) interface{} {
	ret := _m.Called(ctx, post)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*gin.Context, *domain.PostDatum) interface{}); ok {
		r0 = rf(ctx, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// ExecuteTpl provides a mock function with given fields: w, text, data
func (_m *TemplateHandler) ExecuteTpl(w io.Writer, text string, data interface{}) error {
	ret := _m.Called(w, text, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer, string, interface{}) error); ok {
		r0 = rf(w, text, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FuncMap provides a mock function with given fields: ctx, post, cfg
func (_m *TemplateHandler) FuncMap(ctx *gin.Context, post *domain.PostDatum, cfg tpl.TemplateConfig) template.FuncMap {
	ret := _m.Called(ctx, post, cfg)

	var r0 template.FuncMap
	if rf, ok := ret.Get(0).(func(*gin.Context, *domain.PostDatum, tpl.TemplateConfig) template.FuncMap); ok {
		r0 = rf(ctx, post, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(template.FuncMap)
		}
	}

	return r0
}

// GenericFuncMap provides a mock function with given fields:
func (_m *TemplateHandler) GenericFuncMap() template.FuncMap {
	ret := _m.Called()

	var r0 template.FuncMap
	if rf, ok := ret.Get(0).(func() template.FuncMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(template.FuncMap)
		}
	}

	return r0
}

// Prepare provides a mock function with given fields: c
func (_m *TemplateHandler) Prepare(c tpl.TemplateConfig) tpl.TemplateExecutor {
	ret := _m.Called(c)

	var r0 tpl.TemplateExecutor
	if rf, ok := ret.Get(0).(func(tpl.TemplateConfig) tpl.TemplateExecutor); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tpl.TemplateExecutor)
		}
	}

	return r0
}
