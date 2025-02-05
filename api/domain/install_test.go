// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstallVerbis_ToUser(t *testing.T) {
	tt := map[string]struct {
		input *InstallVerbis
		want  interface{}
	}{
		"Success": {
			&InstallVerbis{
				InstallUser: InstallUser{
					UserFirstName:       "Verbis",
					UserLastName:        "CMS",
					UserEmail:           "hello@verbiscms.com",
					UserPassword:        "password",
					UserConfirmPassword: "password",
				},
			},
			UserCreate{
				User: User{
					UserPart: UserPart{
						FirstName: "Verbis",
						LastName:  "CMS",
						Email:     "hello@verbiscms.com",
						Role: Role{
							ID: OwnerRoleID,
						},
					},
				},
				Password:        "password",
				ConfirmPassword: "password",
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.ToUser()
			assert.Equal(t, test.want, got)
		})
	}
}
