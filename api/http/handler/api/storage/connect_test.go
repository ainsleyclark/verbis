// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/gin-gonic/gin"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/services/storage"
	"net/http"
)

func (t *StorageTestSuite) TestStorage_Connect() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Provider)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully updated storage options",
			storageChange,
			func(m *mocks.Provider) {
				m.On("Connect", storageChange).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "provider", Message: "Provider is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			storageChangeBadValidation,
			nil,
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			storageChange,
			func(m *mocks.Provider) {
				m.On("Connect", storageChange).Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			storageChange,
			func(m *mocks.Provider) {
				m.On("Connect", storageChange).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			storageChange,
			func(m *mocks.Provider) {
				m.On("Connect", storageChange).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/storage", "/storage", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Connect(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
