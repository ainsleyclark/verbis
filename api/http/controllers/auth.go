package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
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

// AuthController defines the handler for Authentication methods
type AuthController struct {
	store *models.Store
	config    config.Configuration
}

// newAuth - Construct
func newAuth(m *models.Store, config config.Configuration) *AuthController {
	return &AuthController{
		store: m,
		config:    config,
	}
}

// Login the user
//
// Returns 200 if login was successful.
// Returns 400 if the validation failed.
// Returns 401 if the credentials didn't match.
func (c *AuthController) Login(g *gin.Context) {
	const op = "AuthHandler.Login"

	type login struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var u login
	if err := g.ShouldBindJSON(&u); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	lu, err := c.store.Auth.Authenticate(u.Email, u.Password)
	if err != nil {
		Respond(g, 401, errors.Message(err), err)
		return
	}

	user, err := c.store.User.GetById(lu.Id)
	if err != nil {
		Respond(g, 401, errors.Message(err), err)
		return
	}
	user.Password = ""

	g.SetCookie("verbis-session", user.Token, 172800, "/", "", false, true)

	Respond(g, 200, "Successfully logged in & session started", user)
}

// Logout the user
//
// Returns 200 if logout was successful.
// Returns 400 if the user wasn't found.
// Returns 500 if there was an error logging out.
func (c *AuthController) Logout(g *gin.Context) {
	const op = "AuthHandler.Logout"

	token := g.Request.Header.Get("token")
	fmt.Println(token)
	_, err := c.store.Auth.Logout(token)
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
func (c *AuthController) VerifyEmail(g *gin.Context) {
	const op = "AuthHandler.VerifyEmail"

	token := g.Param("token")
	err := c.store.Auth.VerifyEmail(token)
	if err != nil {
		NoPageFound(g)
		return
	}

	g.Redirect(301, c.config.Admin.Path)
}

// Reset password
//
// Returns 200 if successful.
// Returns 400 if the ID wasn't passed or failed to convert.
func (c *AuthController) ResetPassword(g *gin.Context) {
	const op = "AuthHandler.ResetPassword"

	type resetPassword struct {
		NewPassword     string `json:"new_password" binding:"required,min=8,max=60"`
		ConfirmPassword string `json:"confirm_password" binding:"eqfield=NewPassword,required"`
		Token           string `db:"token" json:"token" binding:"required"`
	}

	var rp resetPassword
	if err := g.ShouldBindJSON(&rp); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if err := c.store.Auth.ResetPassword(rp.Token, rp.NewPassword); err != nil {
		Respond(g, 400, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully reset password", nil)
}

// VerifyPasswordToken
//
// Returns 200 if successful.
// Returns 404 if the token does not exist.
func (c *AuthController) VerifyPasswordToken(g *gin.Context) {
	const op = "AuthHandler.VerifyPasswordToken"

	token := g.Param("token")
	err := c.store.Auth.VerifyPasswordToken(token)
	if err != nil {
		Respond(g, 404, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully logged in & session started", nil)
}

// SendResetPassword reset password email & generate token
//
// Returns 200 if successful.
// Returns 400 if validation failed the user wasn't found.
func (c *AuthController) SendResetPassword(g *gin.Context) {
	const op = "AuthHandler.SendResetPassword"

	type sendResetPassword struct {
		Email string `json:"email" binding:"required,email"`
	}

	var srp sendResetPassword
	if err := g.ShouldBindJSON(&srp); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err := c.store.Auth.SendResetPassword(srp.Email)
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
