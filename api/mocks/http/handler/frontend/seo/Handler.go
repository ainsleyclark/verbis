// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Robots provides a mock function with given fields: ctx
func (_m *Handler) Robots(ctx *gin.Context) {
	_m.Called(ctx)
}

// SiteMapIndex provides a mock function with given fields: ctx
func (_m *Handler) SiteMapIndex(ctx *gin.Context) {
	_m.Called(ctx)
}

// SiteMapResource provides a mock function with given fields: ctx
func (_m *Handler) SiteMapResource(ctx *gin.Context) {
	_m.Called(ctx)
}

// SiteMapXSL provides a mock function with given fields: ctx, index
func (_m *Handler) SiteMapXSL(ctx *gin.Context, index bool) {
	_m.Called(ctx, index)
}
