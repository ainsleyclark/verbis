// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new paths Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for paths to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "paths"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(func() string {return d.Paths.Base},
			"basePath",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(func() string {return d.Paths.Admin},
			"adminPath",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(func() string {return d.Paths.API},
			"apiPath",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(func() string {return d.Paths.Theme},
			"themePath",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(func() string {return d.Paths.Uploads},
			"uploadsPath",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(func() string {return d.Theme.AssetsPath},
			"assetsPath",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(func() string {return d.Paths.Storage},
			"storagePath",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Templates,
			"templatesPath",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Layouts,
			"layoutsPath",
			nil,
			[][2]string{},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
