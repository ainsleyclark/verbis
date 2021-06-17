// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSite_Config(t *testing.T) {
	opts := &domain.Options{
		SiteTitle:       "title",
		SiteDescription: "description",
		SiteLogo:        "logo",
		SiteUrl:         "url",
	}
	want := domain.Site{
		Title:       opts.SiteTitle,
		Description: opts.SiteDescription,
		Logo:        opts.SiteLogo,
		Url:         opts.SiteUrl,
		Version:     version.Version,
	}
	s := New(opts)
	assert.Equal(t, want, s.Global())
}
