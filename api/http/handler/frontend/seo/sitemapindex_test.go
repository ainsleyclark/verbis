// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	mocks "github.com/verbiscms/verbis/api/mocks/publisher"
	"net/http"
)

func (t *SEOTestSuite) TestSEO_SitemapIndex() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		mock    func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context)
	}{
		"Success": {
			testString,
			http.StatusOK,
			"application/xml; charset=utf-8",
			func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context) {
				ms.On("Index").Return(*t.bytes, nil)
			},
		},
		"Fail": {
			testString,
			http.StatusNotFound,
			"text/html",
			func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context) {
				ms.On("Index").Return(nil, fmt.Errorf("error"))
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(http.StatusNotFound, "text/html", []byte(testString))
				})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/sitemap.xml", "/sitemap.xml", nil, func(ctx *gin.Context) {
				t.SetupSitemap(test.mock, ctx).SiteMapIndex(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
