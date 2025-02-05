// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/store/config"
)

// Repository defines methods for redirects
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.Redirects, int, error)
	Find(id int) (domain.Redirect, error)
	FindByFrom(from string) (domain.Redirect, error)
	Create(redirect domain.Redirect) (domain.Redirect, error)
	Update(redirect domain.Redirect) (domain.Redirect, error)
	Delete(id int) error
	Exists(id int) bool
	ExistsByFrom(from string) bool
}

// Store defines the data layer for redirects.
type Store struct {
	*config.Config
}

const (
	// The database table name for redirects.
	TableName = "redirects"
)

var (
	// ErrRedirectExists is returned by validate when
	// a redirect from path already exists.
	ErrRedirectExists = errors.New("redirect already exists")
)

// New
//
// Creates a new redirects store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
