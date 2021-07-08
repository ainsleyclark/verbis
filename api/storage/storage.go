// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/ainsleyclark/verbis/api/store/options"
)

// Provider - TODO
type Provider interface {
	Name() string
	Set(provider domain.StorageProvider) error
}

// Container - TODO
type Container interface {
	List() (domain.Buckets, error)
	Create(name string) error
	Delete(name string) error
}

// Bucket defines the methods used for interacting with
// the Verbis storage system. The client can be remote
// or work with the local file system dependant on
// what is set in the options table.
type Bucket interface {
	// Find looks up the file with the given URL and retrieves
	// the appropriate bucket to obtain the file contents.
	// It returns the byte value of the file as well as
	// the domain.File.
	// Returns errors.INTERNAL if the file could not be opened or read.
	// Returns errors.NOTFOUND if the file could not be retrieved from the bucket.
	Find(url string) ([]byte, domain.File, error)
	// Upload adds a domain.Upload to the database as well as
	// the bucket that is currently set in the env. The
	// file is seeked, the mime type is obtained and it
	// is inserted into the database and uploaded to
	// the bucket.
	// Returns errors.INVALID if the bucket could not be obtained.
	// Returns errors.INTERNAL if the contents couldn't be seeked or the mime
	// type could not be obtained.
	Upload(upload domain.Upload) (domain.File, error)
	// Delete removes an item from the the bucket. It first retrieves
	// the file by a lookup from the database, obtains the correct
	// bucket, then removes the file from the storage provider.
	// The file data will also be deleted from
	// the database.
	// Returns errors.INVALID if the file could not be deleted from the bucket.
	Delete(id int) error
	// Exists queries the database by the given name and
	// returns true if there was a match.
	Exists(name string) bool
}

type Storage struct {
	Provider  Provider
	Container Container
	Bucket    Bucket
}

var (
	// ErrNoProvider is returned by New and SetProvider when
	// there is no match from the options table.
	ErrNoProvider = errors.New("invalid provider")
)

// New parse config
func New(env *environment.Env, opts options.Repository, repo files.Repository) (*Storage, error) {
	const op = "Storage.New"

	c := config{
		Environment: env,
		OptionsRepo: opts,
		FilesRepo:   repo,
		Paths:       paths.Get(),
		Options:     opts.Struct(),
	}

	return &Storage{
		Provider:  &provider{c},
		Container: &container{c},
		Bucket:    &bucket{c},
	}, nil

	//provider := s.opts.StorageProvider
	//if provider == "" {
	//	provider = domain.StorageLocal
	//}
	//
	//if !validate(provider) {
	//	return nil, &errors.Error{Code: errors.INVALID, Message: string("Error setting up storage with provider: " + provider), Operation: op, Err: ErrNoProvider}
	//}

	//err := s.Set(provider)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = s.SetBucket(s.opts.StorageBucket)
	//if err != nil {
	//	return nil, err
	//}
}

// TODO
// validate checks if a string exists in a slice,
func validate(p domain.StorageProvider) bool {
	for _, sp := range domain.StorageProviders {
		if sp == p {
			return true
		}
	}
	return false
}
