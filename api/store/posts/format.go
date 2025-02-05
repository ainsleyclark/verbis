// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"path/filepath"
)

// format traverses the raw posts and constructs out new
// domain.PostData from the results.
func (s *Store) format(raw []postsRaw, layout bool) domain.PostData {
	var posts = make(domain.PostData, 0)

	opts := s.options.Struct()

	cfg, err := s.Theme.Get(opts.ActiveTheme)
	if err != nil {
		logger.WithError(err).Panic()
	}

	for _, v := range raw {
		if !s.find(posts, v.ID) {
			p := domain.PostDatum{
				Post:     v.Post,
				Author:   v.Author.HideCredentials(),
				Fields:   make(domain.PostFields, 0),
				Category: v.Category,
			}

			if layout {
				// TODO, Cacheable is always false.
				theme, err := s.options.GetTheme()
				if err != nil {
					logger.WithError(err).Panic()
				}
				p.Layout = s.finder.Layout(filepath.Join(s.Paths.Themes, theme), p, false)
			}

			p.Permalink = s.permalink(&p, opts, cfg)

			posts = append(posts, p)
		}

		if v.Field.UUID != nil {
			field := domain.PostField{
				PostID:        v.Field.PostId,
				UUID:          *v.Field.UUID,
				Type:          v.Field.Type,
				Name:          v.Field.Name,
				Key:           v.Field.Key,
				Value:         nil,
				OriginalValue: domain.FieldValue(v.Field.OriginalValue),
			}

			for fi, fv := range posts {
				if fv.ID == v.ID {
					posts[fi].Fields = append(posts[fi].Fields, field)
				}
			}
		}
	}

	return posts
}

// find checks if the post data is already in the slice.
func (s *Store) find(posts domain.PostData, id int) bool {
	for _, v := range posts {
		if v.ID == id {
			return true
		}
	}
	return false
}
