// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// SEOHandler is an autogenerated mock type for the SEOHandler type
type SEOHandler struct {
	mock.Mock
}

// Robots provides a mock function with given fields: g
func (_m *SEOHandler) Robots(g *gin.Context) {
	_m.Called(g)
}

// SiteMapIndex provides a mock function with given fields: g
func (_m *SEOHandler) SiteMapIndex(g *gin.Context) {
	_m.Called(g)
}

// SiteMapResource provides a mock function with given fields: g
func (_m *SEOHandler) SiteMapResource(g *gin.Context) {
	_m.Called(g)
}

// SiteMapXSL provides a mock function with given fields: g, index
func (_m *SEOHandler) SiteMapXSL(g *gin.Context, index bool) {
	_m.Called(g, index)
}
