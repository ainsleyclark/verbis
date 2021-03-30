// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
)

// GetTheme
//
// Returns the currently active theme within the options
// table if it has been retrieved successfully.
// Returns errors.NOTFOUND if the 'active_theme' column could not be found.
// Returns errors.INVALID if the option value could not be successfully parsed to a string.
func (s *Store) GetTheme() (string, error) {
	const op = "OptionStore.GetTheme"

	opt, err := s.Find("active_theme")
	if err != nil {
		return "", &errors.Error{Code: errors.NOTFOUND, Message: "No theme exists in the option table: `active_theme`", Operation: op, Err: err}
	}

	theme, err := cast.ToStringE(opt)
	if err != nil {
		fmt.Println(err)
		return "", &errors.Error{Code: errors.INVALID, Message: "Error casting option value to string", Operation: op, Err: err}
	}

	return theme, nil
}

// SetTheme
//
// Returns nil if the theme has been updated successfully.
func (s *Store) SetTheme(theme string) error {
	return s.Update("active_theme", theme)
}
