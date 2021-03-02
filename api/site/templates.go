// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Templates
//
// Retrieves all templates stored within the templates
// directory of the theme path.
//
// Returns errors.NOTFOUND if no templates were found.
// Returns errors.INTERNAL if the template path is invalid.
func (s *Site) Templates(themePath string) (domain.Templates, error) {
	const op = "SiteRepository.Templates"

	tplDir := themePath + s.config.TemplateDir
	files, err := s.walkMatch(tplDir, "*"+s.config.FileExtension)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error getting templates with the path: %s", tplDir), Operation: op, Err: ErrNoTemplates}
	}

	var templates domain.Templates
	for _, file := range files {
		templates = append(templates, domain.Template{
			Key:  file,
			Name: s.fileName(file),
		})
	}

	if len(templates) == 0 {
		return domain.Templates{}, &errors.Error{Code: errors.NOTFOUND, Message: "No page templates available", Operation: op, Err: ErrNoTemplates}
	}

	return templates, nil
}
