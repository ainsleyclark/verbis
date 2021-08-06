// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	domain "github.com/verbiscms/verbis/api/domain"

	io "io"

	mock "github.com/stretchr/testify/mock"

	tpl "github.com/verbiscms/verbis/api/tpl"
)

// TemplateExecutor is an autogenerated mock type for the TemplateExecutor type
type TemplateExecutor struct {
	mock.Mock
}

// Config provides a mock function with given fields:
func (_m *TemplateExecutor) Config() tpl.TemplateConfig {
	ret := _m.Called()

	var r0 tpl.TemplateConfig
	if rf, ok := ret.Get(0).(func() tpl.TemplateConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tpl.TemplateConfig)
		}
	}

	return r0
}

// Execute provides a mock function with given fields: w, name, data
func (_m *TemplateExecutor) Execute(w io.Writer, name string, data interface{}) (string, error) {
	ret := _m.Called(w, name, data)

	var r0 string
	if rf, ok := ret.Get(0).(func(io.Writer, string, interface{}) string); ok {
		r0 = rf(w, name, data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Writer, string, interface{}) error); ok {
		r1 = rf(w, name, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecutePost provides a mock function with given fields: w, name, ctx, post
func (_m *TemplateExecutor) ExecutePost(w io.Writer, name string, ctx *gin.Context, post *domain.PostDatum) (string, error) {
	ret := _m.Called(w, name, ctx, post)

	var r0 string
	if rf, ok := ret.Get(0).(func(io.Writer, string, *gin.Context, *domain.PostDatum) string); ok {
		r0 = rf(w, name, ctx, post)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Writer, string, *gin.Context, *domain.PostDatum) error); ok {
		r1 = rf(w, name, ctx, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Executor provides a mock function with given fields:
func (_m *TemplateExecutor) Executor() tpl.TemplateExecutor {
	ret := _m.Called()

	var r0 tpl.TemplateExecutor
	if rf, ok := ret.Get(0).(func() tpl.TemplateExecutor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tpl.TemplateExecutor)
		}
	}

	return r0
}

// Exists provides a mock function with given fields: template
func (_m *TemplateExecutor) Exists(template string) bool {
	ret := _m.Called(template)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(template)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
