// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"database/sql"
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// Delete
//
// Returns nil if the category was successfully deleted.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if the category was not found.
func (s *Store) Delete(id int) error {
	const op = "CategoryStore.Delete"

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("id", "=", id)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Delete from post meta.
	err = s.meta.Delete(id)
	if err != nil {
		logger.WithError(err).Error()
	}

	// Delete from post fields.
	err = s.fields.Delete(id)
	if err != nil {
		logger.WithError(err).Error()
	}

	// Delete from post categories.
	err = s.categories.Delete(id)
	if err != nil {
		logger.WithError(err).Error()
	}

	return nil
}
