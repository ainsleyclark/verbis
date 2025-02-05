// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/store/users"
	"net/http"
)

func (t *UsersTestSuite) TestUser_Create() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			user,
			http.StatusOK,
			"Successfully created user with ID: 123",
			userCreate,
			func(m *mocks.Repository) {
				m.On("Create", userCreate).Return(user, nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "role_id", Message: "Role ID is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			userCreateBadValidation,
			func(m *mocks.Repository) {
				m.On("Create", userCreateBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			userCreate,
			func(m *mocks.Repository) {
				m.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			userCreate,
			func(m *mocks.Repository) {
				m.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			userCreate,
			func(m *mocks.Repository) {
				m.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/users", "/users", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Create(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
