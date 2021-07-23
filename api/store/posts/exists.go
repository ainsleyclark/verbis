// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// Exists
//
// Returns a bool indicating if the post exists by ID.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(id int) bool {
	const op = "PostStore.Exists"

	q := s.Builder().
		Select("id").
		From(s.Schema()+TableName).
		Where("id", "=", id).
		Exists()

	var exists bool
	err := s.DB().QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}

// ExistsBySlug
//
// Returns a bool indicating if the post exists by slug.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) ExistsBySlug(slug string) bool {
	const op = "PostStore.Exists"

	q := s.Builder().
		Select("id").
		From(s.Schema()+TableName).
		Where("slug", "=", slug).
		Exists()

	var exists bool
	err := s.DB().QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}
