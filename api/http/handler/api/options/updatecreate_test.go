// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	validation "github.com/ainsleyclark/verbis/api/common/vaidation"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/options"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *OptionsTestSuite) TestOptions_UpdateCreate() {
	jsonVOptions, err := json.Marshal(optionsStruct)
	if err != nil {
		t.Error(err)
	}

	dbOptions := domain.OptionsDBMap{}
	err = json.Unmarshal(jsonVOptions, &dbOptions)
	t.NoError(err)

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully created/updated options",
			optionsStruct,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "site_url", Message: "Site Url is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			optionsBadValidation,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
		},
		"Validation Failed DB": {
			nil,
			http.StatusBadRequest,
			"Validation failed",
			"test",
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			optionsStruct,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/posts", "/posts", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).UpdateCreate(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
