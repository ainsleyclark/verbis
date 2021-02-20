// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/gin-gonic/gin"
)

func api(d *deps.Deps, s *server.Server) {

	h := handler.NewApi(d)

	// API Routes
	api := s.Group("/api/v1")
	{
		// Set API Middleware
		api.Use(middleware.EmptyBody())

		// Site
		api.GET("/site", h.Site.Global)

		// Auth
		api.POST("/login", h.Auth.Login)
		api.POST("/logout", h.Auth.Logout)
		api.POST("/password/reset", h.Auth.ResetPassword)
		api.POST("/password/email", h.Auth.SendResetPassword)
		api.GET("/password/verify/:token", h.Auth.VerifyPasswordToken)
		//	api.GET("/email/verify/:token", h.Auth.VerifyEmail)

		// Forms
		forms := api.Group("/forms")
		//forms.Use(csrf.Middleware(csrf.Options{
		//	Secret: "verbis",
		//	ErrorFunc: func(g *gin.Context) {
		//		respond.AbortJSON(g, 400, "CSRF token mismatch", nil)
		//		return
		//	},
		//}))

		forms.POST("/:uuid", h.Forms.Send)

		// Operator
		operator := api.Group("")
		{
			operator.Use(middleware.OperatorTokenCheck(d))
			operator.Use(middleware.SessionCheck(d))

			// Theme
			operator.GET("/theme", h.Site.Theme)

			// Templates
			operator.GET("/templates", h.Site.Templates)

			// Layouts
			operator.GET("/layouts", h.Site.Layouts)

			// Posts
			operator.GET("/posts", h.Posts.List)
			operator.GET("/posts/:id", h.Posts.Find)
			operator.POST("/posts", h.Posts.Create)
			operator.PUT("/posts/:id", h.Posts.Update)
			operator.DELETE("/posts/:id", h.Posts.Delete)

			// Categories
			operator.GET("/categories", h.Categories.List)
			operator.GET("/categories/:id", h.Categories.Find)
			operator.POST("/categories", h.Categories.Create)
			operator.PUT("/categories/:id", h.Categories.Update)
			operator.DELETE("/categories/:id", h.Categories.Delete)

			// Media
			operator.GET("/media", h.Media.List)
			operator.GET("/media/:id", h.Media.Find)
			operator.POST("/media", h.Media.Upload)
			operator.PUT("/media/:id", h.Media.Update)
			operator.DELETE("/media/:id", h.Media.Delete)

			// Users
			operator.GET("/users", h.Users.List)
			operator.GET("/users/:id", h.Users.Find)
			operator.PUT("/users/:id", h.Users.Update)
			operator.POST("/users/:id/reset-password", h.Users.ResetPassword)

			// Fields
			operator.GET("/fields", h.Fields.List)

			// Options
			operator.GET("/options", h.Options.List)
			operator.GET("/options/:name", h.Options.Find)
			operator.POST("/options", h.Options.UpdateCreate)

			// Roles
			operator.GET("/roles", h.Users.Roles)

			// Redirects
			operator.GET("/redirects", h.Redirects.List)
			operator.GET("/redirects/:id", h.Redirects.Find)
			operator.POST("/redirects", h.Redirects.Create)
			operator.PUT("/redirects/:id", h.Redirects.Update)
			operator.DELETE("/redirects/:id", h.Redirects.Delete)

			// Cache
			operator.POST("/cache", h.Cache.Clear)
		}

		// Administrator
		admin := api.Group("")
		{
			admin.Use(middleware.AdminTokenCheck(d))
			operator.Use(middleware.SessionCheck(d))

			// Users
			admin.POST("/users", h.Users.Create)
			admin.DELETE("/users/:id", h.Users.Delete)
		}
	}

	api.Use(gin.Recovery())
}
