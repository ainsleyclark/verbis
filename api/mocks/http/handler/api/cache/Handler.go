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

// Clear provides a mock function with given fields: ctx
func (_m *Handler) Clear(ctx *gin.Context) {
	_m.Called(ctx)
}
