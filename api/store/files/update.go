// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Update
//
// Returns a new file upon successful update.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Update(f domain.File) (domain.File, error) {
	const op = "FileStore.Update"

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("uuid", f.UUID.String()).
		Column("url", f.Url).
		Column("name", f.Name).
		Column("bucket_id", f.BucketId).
		Column("mime", f.Mime).
		Column("source_type", f.SourceType).
		Column("provider", f.Provider).
		Column("region", f.Region).
		Column("bucket", f.Bucket).
		Column("file_size", f.FileSize).
		Column("private", f.Private).
		Where("id", "=", f.Id)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating file with the name: " + f.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return f, nil
}
