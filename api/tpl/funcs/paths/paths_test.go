// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"testing"
)

var (
	ns = New(&deps.Deps{
		Paths: paths.Paths{
			Base:    "/base",
			Admin:   "/admin/",
			API:     "/api/",
			Uploads: "/uploads/",
			Storage: "/storage/",
			Web:     "/web/",
		},
		Config: &domain.ThemeConfig{
			TemplateDir: "templates",
			LayoutDir:   "layouts",
			AssetsPath:  "/assets/",
		},
		Options: &domain.Options{
			ActiveTheme: "theme",
		},
	})
)

func TestNamespace_Base(t *testing.T) {
	got := ns.Base()
	assert.Equal(t, "/base", got)
}

func TestNamespace_Admin(t *testing.T) {
	got := ns.Admin()
	assert.Equal(t, "/admin/", got)
}

func TestNamespace_API(t *testing.T) {
	got := ns.API()
	assert.Equal(t, "/api/", got)
}

func TestNamespace_Theme(t *testing.T) {
	got := ns.Theme()
	assert.Equal(t, "/base/themes/theme", got)
}

func TestNamespace_Uploads(t *testing.T) {
	got := ns.Uploads()
	assert.Equal(t, "/uploads/", got)
}

func TestNamespace_Storage(t *testing.T) {
	got := ns.Storage()
	assert.Equal(t, "/storage/", got)
}

func TestNamespace_Assets(t *testing.T) {
	got := ns.Assets()
	assert.Equal(t, "/assets/", got)
}

func TestNamespace_Templates(t *testing.T) {
	got := ns.Templates()
	assert.Equal(t, "/base/themes/theme/templates", got)
}

func TestNameSpace_Layouts(t *testing.T) {
	got := ns.Layouts()
	assert.Equal(t, "/base/themes/theme/layouts", got)
}
