// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

func (t *ThemeTestSuite) TestSite_Templates() {
	tt := map[string]struct {
		theme string
		want  interface{}
	}{
		"Success": {
			"verbis",
			domain.Templates{
				{Key: "nested/template-nested", Name: "Nested/template Nested"},
				{Key: "template-hyphen", Name: "template Hyphen"},
				{Key: "template", Name: "template"},
			},
		},
		"Wrong Path": {
			"wrong",
			"Error getting templates with the path:",
		},
		"No Layouts": {
			"empty",
			"No templates available",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup()
			got, err := s.Templates(test.theme)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
