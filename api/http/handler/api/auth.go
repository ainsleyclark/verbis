// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
)

// AuthHandler defines methods for auth methods to interact with the server
type AuthHandler interface {
	Login(g *gin.Context)
	Logout(g *gin.Context)
	ResetPassword(g *gin.Context)
	VerifyEmail(g *gin.Context)
	VerifyPasswordToken(g *gin.Context)
	SendResetPassword(g *gin.Context)
}

// Auth defines the handler for Authentication methods
type Auth struct {
	*deps.Deps
}

// newAuth - Construct
func NewAuth(d *deps.Deps) *Auth {
	return &Auth{d}
}

// Login the user
//
// Returns 200 if login was successful.
// Returns 400 if the validation failed.
// Returns 401 if the credentials didn't match.
func (c *Auth) Login(g *gin.Context) {
	const op = "AuthHandler.Login"

	var l domain.Login
	if err := g.ShouldBindJSON(&l); err != nil {
		fmt.Print(err)
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := c.Store.Auth.Authenticate(l.Email, l.Password)
	if err != nil {
		fmt.Print(err)
		Respond(g, 401, errors.Message(err), err)
		return
	}
	user.HidePassword()

	g.SetCookie("verbis-session", user.Token, 172800, "/", "", false, true)

	Respond(g, 200, "Successfully logged in & session started", user)
}

// Logout the user
//
// Returns 200 if logout was successful.
// Returns 400 if the user wasn't found.
// Returns 500 if there was an error logging out.
func (c *Auth) Logout(g *gin.Context) {
	const op = "AuthHandler.Logout"

	token := g.Request.Header.Get("token")
	_, err := c.Store.Auth.Logout(token)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	g.SetCookie("verbis-session", "", -1, "/", "", false, true)

	Respond(g, 200, "Successfully logged out", nil)
}

// Verify email
//
// TODO
func (c *Auth) VerifyEmail(g *gin.Context) {
	const op = "AuthHandler.VerifyEmail"

	token := g.Param("token")
	err := c.Store.Auth.VerifyEmail(token)
	if err != nil {
		notFound(g)
		return
	}

	g.Redirect(301, c.Store.Config.Admin.Path)
}

// Reset password
//
// Returns 200 if successful.
// Returns 400 if the ID wasn't passed or failed to convert.
func (c *Auth) ResetPassword(g *gin.Context) {
	const op = "AuthHandler.ResetPassword"

	var rp domain.ResetPassword
	if err := g.ShouldBindJSON(&rp); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err := c.Store.Auth.ResetPassword(rp.Token, rp.NewPassword)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully reset password", nil)
}

// VerifyPasswordToken
//
// Returns 200 if successful.
// Returns 404 if the token does not exist.
func (c *Auth) VerifyPasswordToken(g *gin.Context) {
	const op = "AuthHandler.VerifyPasswordToken"

	err := c.Store.Auth.VerifyPasswordToken(g.Param("token"))
	if err != nil {
		Respond(g, 404, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully verified token", nil)
}

// SendResetPassword reset password email & generate token
//
// Returns 200 if successful.
// Returns 400 if validation failed the user wasn't found.
func (c *Auth) SendResetPassword(g *gin.Context) {
	const op = "AuthHandler.SendResetPassword"

	var srp domain.SendResetPassword
	if err := g.ShouldBindJSON(&srp); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err := c.Store.Auth.SendResetPassword(srp.Email)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	}
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "A fresh verification link has been sent to your email", nil)
}
