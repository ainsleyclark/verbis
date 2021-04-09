// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"fmt"
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/errors"
)

// ResetPassword defines the event instance for resetting passwords
type ResetPassword struct {
	ev *Event
}

// Reset Password
//
//
func NewResetPassword(e *Event) *ResetPassword {
	e.subject = "Verbis - Reset Password"
	e.template = "reset-password"
	e.plaintext = ""
	return &ResetPassword{
		ev: e,
	}
}

// Dispatch
//
//
func (r *ResetPassword) Dispatch(d Data, recipients []string, attachments mail.Attachments) error {
	err := r.Validate(d)
	if err != nil {
		return err
	}

	site := r.ev.Deps.Site.Global()
	d["Url"] = site.Url + "/password/reset/" + d["Token"].(string)

	err = r.ev.send(d, recipients, attachments)
	if err != nil {
		return err
	}

	return nil
}

// Validate
//
//
func (r *ResetPassword) Validate(d Data) error {
	const op = "Events.ResetPassword.Validate"

	if !d.Exists("Token") {
		return &errors.Error{Code: errors.INVALID, Message: "Token cannot be empty to send reset password event.", Operation: op, Err: fmt.Errorf("token must not be empty")}
	}

	if !d.Exists("User") {
		return &errors.Error{Code: errors.INVALID, Message: "User cannot be empty to send reset password event.", Operation: op, Err: fmt.Errorf("user must not be empty")}
	}

	_, ok := d["Token"].(string)
	if !ok {
		return &errors.Error{Code: errors.INVALID, Message: "Token must be a string", Operation: op, Err: fmt.Errorf("token must be a string")}
	}

	return nil
}
