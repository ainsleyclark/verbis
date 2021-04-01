// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *PostsTestSuite) TestPosts_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			postData,
			http.StatusOK,
			"Successfully updated post with ID: 123",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(postData, nil)
			},
			"/posts/123",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "post_title", Message: "Post Title is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			postBadValidation,
			func(m *mocks.Repository) {
				m.On("Update", postBadValidation).Return(domain.PostDatum{}, fmt.Errorf("error"))
			},
			"/posts/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to update the post",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(domain.PostDatum{}, fmt.Errorf("error"))
			},
			"/posts/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/posts/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"config",
			post,
			func(m *mocks.Repository) {
				m.On("Update", postCreate).Return(domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: "config"})
			},
			"/posts/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPut, test.url, "/posts/:id", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
