// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Config provides a mock function with given fields: ctx
func (_m *Handler) Config(ctx *gin.Context) {
	_m.Called(ctx)
}

// CreateBucket provides a mock function with given fields: ctx
func (_m *Handler) CreateBucket(ctx *gin.Context) {
	_m.Called(ctx)
}

// DeleteBucket provides a mock function with given fields: ctx
func (_m *Handler) DeleteBucket(ctx *gin.Context) {
	_m.Called(ctx)
}

// ListBuckets provides a mock function with given fields: ctx
func (_m *Handler) ListBuckets(ctx *gin.Context) {
	_m.Called(ctx)
}

// Migrate provides a mock function with given fields: ctx
func (_m *Handler) Migrate(ctx *gin.Context) {
	_m.Called(ctx)
}

// Save provides a mock function with given fields: ctx
func (_m *Handler) Save(ctx *gin.Context) {
	_m.Called(ctx)
}
