// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	mocks "github.com/verbiscms/verbis/api/mocks/publisher"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// SEOTestSuite defines the helper used for seo
// testing.
type SEOTestSuite struct {
	test.HandlerSuite
	bytes *[]byte
}

// TestSEO
//
// Assert testing has begun.
func TestSEO(t *testing.T) {
	b := []byte(testString)
	suite.Run(t, &SEOTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
		bytes:        &b,
	})
}

// Setup
//
// A helper to obtain a seo handler for testing.
func (t *SEOTestSuite) Setup(mf func(ms *mocks.Publisher, ctx *gin.Context), ctx *gin.Context) *SEO {
	m := &mocks.Publisher{}
	if mf != nil {
		mf(m, ctx)
	}
	return &SEO{
		Deps:      &deps.Deps{},
		publisher: m,
	}
}

// SetupSitemap
//
// A helper to obtain a seo handler for testing
// for sitemap handle funcs.
func (t *SEOTestSuite) SetupSitemap(mf func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context), ctx *gin.Context) *SEO {
	ms := &mocks.SiteMapper{}
	m := &mocks.Publisher{}
	m.On("SiteMap").Return(ms)
	if mf != nil {
		mf(m, ms, ctx)
	}
	return &SEO{
		Deps:      &deps.Deps{},
		publisher: m,
	}
}

var (
	// The default string used for testing.
	testString = "test"
)
