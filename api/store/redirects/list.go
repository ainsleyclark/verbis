// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// List
//
// Returns a slice of redirects with the total amount.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no categories available.
func (s *Store) List(meta params.Params) (domain.Redirects, int, error) {
	const op = "RedirectStore.List"

	q := s.Builder().
		From(s.Schema() + TableName)

	// Apply filters.
	err := database.FilterRows(s.Driver, q, meta.Filters, TableName)
	if err != nil {
		return nil, -1, err
	}

	// Apply order.
	if meta.OrderBy != "" {
		q.OrderBy(meta.OrderBy, meta.OrderDirection)
	}
	countQ := q.Count()

	// Apply pagination.
	if !meta.LimitAll {
		q.Limit(meta.Limit).Offset((meta.Page - 1) * meta.Limit)
	}

	// Select redirects.
	var redirects domain.Redirects
	err = s.DB().Select(&redirects, q.Build())
	if err == sql.ErrNoRows {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No redirects available", Operation: op, Err: err}
	} else if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	// Count the total number of redirects.
	var total int
	err = s.DB().QueryRow(countQ).Scan(&total)
	if err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the total number of redirects", Operation: op, Err: err}
	}

	return redirects, total, nil
}
