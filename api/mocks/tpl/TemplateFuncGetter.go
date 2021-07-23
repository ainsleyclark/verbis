// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	domain "github.com/verbiscms/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"

	template "html/template"

	tpl "github.com/verbiscms/verbis/api/tpl"
)

// TemplateFuncGetter is an autogenerated mock type for the TemplateFuncGetter type
type TemplateFuncGetter struct {
	mock.Mock
}

// FuncMap provides a mock function with given fields: ctx, post, cfg
func (_m *TemplateFuncGetter) FuncMap(ctx *gin.Context, post *domain.PostDatum, cfg tpl.TemplateConfig) template.FuncMap {
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
func (_m *TemplateFuncGetter) GenericFuncMap() template.FuncMap {
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
