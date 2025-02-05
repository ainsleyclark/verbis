// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"github.com/verbiscms/verbis/api/http/pagination"
	"net/http"
)

// List
//
// Returns http.StatusOK if there are no forms or success.
// Returns http.StatusInternalServerError if there was an error getting the forms.
// Returns http.StatusBadRequest if there was conflict or the request was invalid.
func (f *Forms) List(ctx *gin.Context) {
	const op = "FormHandler.List"

	p := api.Params(ctx).Get()

	forms, total, err := f.Store.Forms.List(p)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained forms", forms, pagination.Get(p, total))
}
