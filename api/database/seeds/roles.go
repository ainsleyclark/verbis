// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seeds

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// runRoles will insert all default values for the user roles
// including the Id.
func (s *Seeder) runRoles() error {
	r := []domain.UserRole{
		{
			Id:          1,
			Name:        "Banned",
			Description: "The user has been banned from the system.",
		},
		{
			Id:          2,
			Name:        "Contributor",
			Description: "The user can create and edit their own draft posts, but they are unable to edit drafts of users or published posts.",
		},
		{
			Id:          3,
			Name:        "Author",
			Description: "The user can write, edit and publish their own posts.",
		},
		{
			Id:          4,
			Name:        "Editor",
			Description: "The user can do everything defined in the Author role but they can also edit and publish posts of others, as well as their own.",
		},
		{
			Id:          5,
			Name:        "Administrator",
			Description: "The user can do everything defined in the Editor role but they can also edit site settings and data. Additionally they can manage users",
		},
		{
			Id:          6,
			Name:        "Owner",
			Description: "The user is a special user with all of the permissions as an Administrator however they cannot be deleted",
		},
	}

	for _, v := range r {
		if exists := s.models.Roles.Exists(v.Name); !exists {
			if _, err := s.models.Roles.Create(&v); err != nil {
				return err
			}
		}
	}

	return nil
}
