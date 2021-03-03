// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Templates
//
// Returns http.StatusInternalServerError if there was an error getting the templates.
// Returns http.StatusOK if the templates were obtained successfully or there were none found.
func (s *Site) Templates(ctx *gin.Context) {
	const op = "SiteHandler.Templates"

	templates, err := s.Site.Templates(s.ThemePath())
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained templates", templates)
}
