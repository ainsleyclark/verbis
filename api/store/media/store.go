// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/store/config"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/ainsleyclark/verbis/api/store/media/sizes"
)

// Repository defines methods for media items
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.MediaItems, int, error)
	Find(id int) (domain.Media, error)
	Create(m domain.Media) (domain.Media, error)
	Update(m domain.Media) (domain.Media, error)
	Delete(id int) error
}

// Store defines the data layer for media.
type Store struct {
	*config.Config
	sizes sizes.Repository
}

const (
	// TableName is the database table name for media.
	TableName = "media"
)

// New
//
// Creates a new media store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
		sizes:  sizes.New(cfg),
	}
}

// selectStmt is a helper for SELECT Statements,
// joining files by file id.
func (s *Store) selectStmt() *builder.Sqlbuilder {
	return s.Builder().
		SelectRaw(s.Schema()+TableName+".*, "+
			s.Schema()+"file.id `file.id`, "+
			s.Schema()+"file.url `file.url`, "+
			s.Schema()+"file.name `file.name`, "+
			s.Schema()+"file.path `file.path`, "+
			s.Schema()+"file.mime `file.mime`, "+
			s.Schema()+"file.source_type `file.source_type`, "+
			s.Schema()+"file.provider `file.provider`, "+
			s.Schema()+"file.region `file.region`, "+
			s.Schema()+"file.bucket `file.bucket`, "+
			s.Schema()+"file.file_size `file.file_size`, "+
			s.Schema()+"file.private `file.private`").
		From(s.Schema()+TableName).
		LeftJoin(s.Schema()+files.TableName, "file", s.Schema()+TableName+".file_id = "+s.Schema()+"file.id")
}
