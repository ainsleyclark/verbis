// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import "github.com/verbiscms/verbis/api/errors"

func (t *ConfigTestSuite) Test_FindScreenshot() {
	tt := map[string]struct {
		path string
		want interface{}
	}{
		"PNG": {
			t.TestPath,
			"/themes/testdata/screenshot.png",
		},
		"Wrong Path": {
			"wrong",
			"No screenshot found from the theme",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := findScreenshot(test.path)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
