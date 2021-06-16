// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"os"
	"path/filepath"
)

// Paths represent the struct of paths for use with the
/// application.
type Paths struct {
	Base    string
	Admin   string
	API     string
	Uploads string
	Storage string
	Themes  string
	Web     string
	Forms   string
	Bin     string
}

const (
	Admin   = string(os.PathSeparator) + "admin"
	API     = string(os.PathSeparator) + "api"
	Storage = string(os.PathSeparator) + "storage"
	Themes  = string(os.PathSeparator) + "themes"
	Web     = API + string(os.PathSeparator) + "www"
	Uploads = Storage + string(os.PathSeparator) + "uploads"
	Forms   = Storage + string(os.PathSeparator) + "forms"
	Bin     = string(os.PathSeparator) + "bin"
)

// Get
//
// Retrieves relevant paths for the application.
func Get() Paths {
	base := base()

	return Paths{
		Base:    base,
		Admin:   base + Admin,
		API:     base + API,
		Uploads: base + Uploads,
		Storage: base + Storage,
		Themes:  base + Themes,
		Web:     base + Web,
		Forms:   base + Forms,
		Bin:     base + Bin,
	}
}

// base
//
// Base path of project.
func base() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

// BaseCheck
//
// Check the environment to see if it is passable by
// seeing if the .env file, the admin folder, and
// the storage folder exists.
func BaseCheck() error {
	const op = "paths.BaseCheck"

	basePath := base()

	if !files.Exists(basePath + "/.env") {
		return fmt.Errorf("could not locate the .env file in the current directory")
	}

	storage := basePath + Storage
	if !files.DirectoryExists(storage) {
		err := os.Mkdir(storage, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
