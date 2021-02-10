// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	template "html/template"

	tpl "github.com/ainsleyclark/verbis/api/tpl"
)

// TemplateHandler is an autogenerated mock type for the TemplateHandler type
type TemplateHandler struct {
	mock.Mock
}

// Data provides a mock function with given fields: ctx, post
func (_m *TemplateHandler) Data(ctx *gin.Context, post *domain.PostData) interface{} {
	ret := _m.Called(ctx, post)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*gin.Context, *domain.PostData) interface{}); ok {
		r0 = rf(ctx, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// FuncMap provides a mock function with given fields: ctx, post, cfg
func (_m *TemplateHandler) FuncMap(ctx *gin.Context, post *domain.PostData, cfg tpl.TemplateConfig) template.FuncMap {
	ret := _m.Called(ctx, post, cfg)

	var r0 template.FuncMap
	if rf, ok := ret.Get(0).(func(*gin.Context, *domain.PostData, tpl.TemplateConfig) template.FuncMap); ok {
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
