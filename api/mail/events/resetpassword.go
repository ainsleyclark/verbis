// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mail"
)

// ResetPassword defines the event instance for resetting passwords
type ResetPassword struct {
	mailer *mail.Mail
}

// NewResetPassword creates a new reset password event.
func NewResetPassword() (*ResetPassword, error) {
	const op = "events.NewResetPassword"

	m, err := mail.New()
	if err != nil {
		return &ResetPassword{}, err
	}

	return &ResetPassword{
		mailer: m,
	}, nil
}

// Send the reset password event.
func (e *ResetPassword) Send(u *domain.User, url, token, title string) error {
	const op = "events.ResetPassword.Send"

	data := mail.Data{
		"AppUrl":    url,
		"AppTitle":  title,
		"AdminPath": "/admin",
		"Token":     token,
		"UserName":  u.FirstName,
	}

	html, err := e.mailer.ExecuteHTML("reset-password.html", data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tm := mail.Sender{
		To:      []string{u.Email},
		Subject: "Reset password",
		HTML:    html,
	}

	e.mailer.Send(&tm)

	return nil
}
