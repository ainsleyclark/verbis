// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	pkgValidate "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getUserMock is a helper to obtain a mock user controller
// for testing.
func getUserMock(m models.UserRepository) *User {
	return &User{
		store: &models.Store{
			User: m,
		},
	}
}

// mockComparePassword for the password field on the domain.UserPasswordReset
// (custom validation)
func mockComparePassword(fl pkgValidate.FieldLevel) bool {
	return true
}

// Test_NewUser - Test construct
func Test_NewUser(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &User{
		store:  &store,
		config: config,
	}
	got := NewUser(&store, config)
	assert.Equal(t, got, want)
}

// TestUser_Get - Test Get route
func TestUser_Get(t *testing.T) {

	users := domain.Users{
		{
			UserPart: domain.UserPart{
				Id: 123, FirstName: "Verbis", LastName: "CMS",
			},
		},
		{
			UserPart: domain.UserPart{
				Id: 123, FirstName: "Verbis", LastName: "CMS",
			},
		},
	}

	pagination := params.Params{Page: 1, Limit: 15, OrderBy: "id", OrderDirection: "ASC", Filters: nil}

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		mock    func(u *mocks.UserRepository)
	}{
		"Success": {
			want:    `[{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"},{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}]`,
			status:  200,
			message: "Successfully obtained users",
			mock: func(u *mocks.UserRepository) {
				u.On("Get", pagination).Return(users, 1, nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "no users found",
			mock: func(u *mocks.UserRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"})
			},
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			mock: func(u *mocks.UserRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			want:    `{}`,
			status:  400,
			message: "invalid",
			mock: func(u *mocks.UserRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(u *mocks.UserRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.UserRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/users", "/users", nil, func(g *gin.Context) {
				getUserMock(mock).Get(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestUser_GetById - Test GetByID route
func TestUser_GetById(t *testing.T) {

	user := domain.User{
		UserPart: domain.UserPart{
			Id: 123, FirstName: "Verbis", LastName: "CMS",
		},
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(u *mocks.UserRepository)
		url     string
	}{
		"Success": {
			want:    `{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully obtained user with ID: 123",
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(user, nil)
			},
			url: "/users/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "Pass a valid number to obtain the user by ID",
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, fmt.Errorf("error"))
			},
			url: "/users/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "no users found",
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"})
			},
			url: "/users/123",
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/users/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.UserRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", test.url, "/users/:id", nil, func(g *gin.Context) {
				getUserMock(mock).GetById(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestUser_GetRoles - Test GetRoles route
func TestUser_GetRoles(t *testing.T) {

	roles := []domain.UserRole{
		{Id: 1, Name: "Banned", Description: "Banned Role"},
		{Id: 2, Name: "Administrator", Description: "Administrator Role"},
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(u *mocks.UserRepository)
	}{
		"Success": {
			want:    `[{"description":"Banned Role","id":1,"name":"Banned"},{"description":"Administrator Role","id":2,"name":"Administrator"}]`,
			status:  200,
			message: "Successfully obtained user roles",
			mock: func(u *mocks.UserRepository) {
				u.On("GetRoles").Return(roles, nil)
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(u *mocks.UserRepository) {
				u.On("GetRoles").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.UserRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/roles", "/roles", nil, func(g *gin.Context) {
				getUserMock(mock).GetRoles(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestUser_Create - Test Create route
func TestUser_Create(t *testing.T) {

	userCreate := &domain.UserCreate{
		User: domain.User{
			UserPart: domain.UserPart{
				FirstName: "Verbis",
				LastName:  "CMS",
				Email:     "verbis@verbiscms.com",
				Role: domain.UserRole{
					Id: 123,
				},
			},
		},
		Password:        "password",
		ConfirmPassword: "password",
	}

	user := domain.User{
		UserPart: domain.UserPart{
			Id:        123,
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
		},
	}

	userBadValidation := &domain.UserCreate{
		User: domain.User{
			UserPart: domain.UserPart{
				FirstName: "Verbis",
				LastName:  "CMS",
				Email:     "verbis@verbiscms.com",
			},
		},
		Password:        "password",
		ConfirmPassword: "password",
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		input   interface{}
		mock    func(u *mocks.UserRepository)
	}{
		"Success": {
			want:    `{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"verbis@verbiscms.com","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully created user with ID: 123",
			input:   userCreate,
			mock: func(u *mocks.UserRepository) {
				u.On("Create", userCreate).Return(user, nil)
			},
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"role_id","message":"Role Id is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   userBadValidation,
			mock: func(u *mocks.UserRepository) {
				u.On("Create", userBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			want:    `{}`,
			status:  400,
			message: "invalid",
			input:   userCreate,
			mock: func(u *mocks.UserRepository) {
				u.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			input:   userCreate,
			mock: func(u *mocks.UserRepository) {
				u.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   userCreate,
			mock: func(u *mocks.UserRepository) {
				u.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.UserRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/users", "/users", bytes.NewBuffer(body), func(g *gin.Context) {
				getUserMock(mock).Create(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

//TestUser_Update - Test Update route
func TestUser_Update(t *testing.T) {

	cache.Init()

	user := domain.User{
		UserPart: domain.UserPart{
			Id:        123,
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
			Role: domain.UserRole{
				Id: 1,
			},
		},
	}

	userBadValidation := &domain.User{
		UserPart: domain.UserPart{
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
		},
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		input   interface{}
		mock    func(u *mocks.UserRepository)
		url     string
	}{
		"Success": {
			want:    `{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"verbis@verbiscms.com","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":1,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully updated user with ID: 123",
			input:   user,
			mock: func(u *mocks.UserRepository) {
				u.On("Update", &user).Return(user, nil)
			},
			url: "/users/123",
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"role_id","message":"Role Id is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   userBadValidation,
			mock: func(u *mocks.UserRepository) {
				u.On("Update", userBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
			url: "/users/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "A valid ID is required to update the user",
			input:   user,
			mock: func(u *mocks.UserRepository) {
				u.On("Update", userBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
			url: "/users/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			input:   user,
			mock: func(u *mocks.UserRepository) {
				u.On("Update", &user).Return(domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/users/123",
		},
		"Internal": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   user,
			mock: func(u *mocks.UserRepository) {
				u.On("Update", &user).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/users/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.UserRepository{}
			postsMock := &mocks.PostsRepository{}
			postsMock.On("Get", params.Params{LimitAll: true}, false, "", "").Return([]domain.PostData{}, 2, nil)

			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("PUT", test.url, "/users/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				t := getUserMock(mock)
				t.store.Posts = postsMock
				t.Update(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestUser_Delete - Test Delete route
func TestUser_Delete(t *testing.T) {

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(u *mocks.UserRepository)
		url     string
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully deleted user with ID: 123",
			mock: func(u *mocks.UserRepository) {
				u.On("Delete", 123).Return(nil)
			},
			url: "/users/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "A valid ID is required to delete a user",
			mock: func(u *mocks.UserRepository) {
				u.On("Delete", 123).Return(nil)
			},
			url: "/users/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			mock: func(u *mocks.UserRepository) {
				u.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/users/123",
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			mock: func(u *mocks.UserRepository) {
				u.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			url: "/users/123",
		},
		"Internal": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(u *mocks.UserRepository) {
				u.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/users/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.UserRepository{}
			test.mock(mock)

			rr.RequestAndServe("DELETE", test.url, "/users/:id", nil, func(g *gin.Context) {
				getUserMock(mock).Delete(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestUser_ResetPassword - Test Reset Password route
func TestUser_ResetPassword(t *testing.T) {

	reset := domain.UserPasswordReset{
		DBPassword:      "",
		CurrentPassword: "password",
		NewPassword:     "verbiscms",
		ConfirmPassword: "verbiscms",
	}

	resetBadValidation := &domain.UserPasswordReset{
		CurrentPassword: "password",
		NewPassword:     "verbiscms",
		ConfirmPassword: "verbiscmss",
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		input   interface{}
		mock    func(u *mocks.UserRepository)
		url     string
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully updated password for the user with ID: 123",
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, nil)
				u.On("ResetPassword", 123, reset).Return(nil)
			},
			input: reset,
			url:   "/users/reset/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "A valid ID is required to update a user's password",
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, nil)
				u.On("ResetPassword", 123, reset).Return(nil)
			},
			input: reset,
			url:   "/users/reset/wrongid",
		},
		"Not found": {
			want:    `{}`,
			status:  400,
			message: "No user has been found with the ID: 123",
			input:   reset,
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
				u.On("ResetPassword", 123, reset).Return(nil)
			},
			url: "/users/reset/123",
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"confirm_password", "message":"Confirm Password must equal the New Password.", "type":"eqfield"}]}`,
			status:  400,
			message: "Validation failed",
			input:   resetBadValidation,
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, nil)
				u.On("ResetPassword", 123, reset).Return(nil)
			},
			url: "/users/reset/123",
		},
		"Invalid": {
			want:    `{}`,
			status:  400,
			message: "invalid",
			input:   reset,
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, nil)
				u.On("ResetPassword", 123, reset).Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
			url: "/users/reset/123",
		},
		"Internal": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   reset,
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, nil)
				u.On("ResetPassword", 123, reset).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/users/reset/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.UserRepository{}
			test.mock(mock)

			if v, ok := binding.Validator.Engine().(*pkgValidate.Validate); ok {
				v.RegisterValidation("password", mockComparePassword)
			}

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("DELETE", test.url, "/users/reset/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				getUserMock(mock).ResetPassword(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}
