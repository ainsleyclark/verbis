// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/server"
)

// Vue (SPA) routes
func spa(s *server.Server, c *handler.Handler, config config.Configuration) {
	spa := s.Group(config.Admin.Path)
	{
		spa.GET("/*any", c.SPA.Serve)
		spa.GET("", c.SPA.Serve)
	}
}
