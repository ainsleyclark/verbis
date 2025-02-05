// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/test"
	"net/http"
	"testing"
)

// APITestSuite defines the helper used for api
// testing.
type APITestSuite struct {
	test.HandlerSuite
}

// TestApi
//
// Assert testing has begun.
func TestApi(t *testing.T) {
	suite.Run(t, &APITestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

func (t *APITestSuite) TestParams() {
	ctx := &gin.Context{}
	want := &params.Params{
		Stringer: &apiParams{ctx: ctx},
	}
	got := Params(ctx)
	t.Equal(want.Stringer, got.Stringer)
}

func (t *APITestSuite) TestParam() {
	ctx := &gin.Context{}
	req, err := http.NewRequest(http.MethodGet, "https://verbiscms.com?query=test", nil)
	t.NoError(err)
	ctx.Request = req
	api := &apiParams{ctx: ctx}
	got := api.Param("query")
	t.Equal("test", got)
}

func (t *APITestSuite) TestApiParams_Param() {
	tt := map[string]struct {
		query string
		param string
		want  string
	}{
		"Page": {
			"page=2",
			"page",
			"2",
		},
		"Limit": {
			"limit=2",
			"limit",
			"2",
		},
		"Limit All": {
			"limit=all",
			"limit",
			"all",
		},
		"Order By": {
			"order_by=id",
			"order_by",
			"id",
		},
		"Order Direction": {
			"order_direction=name",
			"order_direction",
			"name",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/test?"+test.query, "/login", nil, func(ctx *gin.Context) {
				got := Params(ctx).Param(test.param)
				t.Equal(test.want, got)
			})
			t.Reset()
		})
	}
}
