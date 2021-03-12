// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postfields

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Exists
//
// Returns a bool indicating if the field exists.
// Logs errors.INTERNAL if there was an error executing the query.
func (s *Store) Exists(field domain.PostField) bool {
	const op = "PostFieldStore.Exists"

	q := s.Builder().
		Select("id").
		From(s.Schema()+TableName).
		Where("post_id", "=", field.Id).
		Where("uuid", "=", field.UUID).
		Where("key", "=", field.Key).
		Where("name", "=", field.Name).
		Exists()

	var exists bool
	err := s.DB().QueryRow(q).Scan(&exists)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}).Error()
	}

	return exists
}
