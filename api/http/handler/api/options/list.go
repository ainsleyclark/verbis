// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// List
//
// Returns http.StatusOK if there are no options or success.
// Returns http.StatusInternalServerError if there was an error getting the options.
func (o *Options) List(ctx *gin.Context) {
	const op = "OptionsHandler.List"

	options, err := o.Store.Options.Map()
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

	api.Respond(ctx, http.StatusOK, "Successfully obtained options", options)
}
