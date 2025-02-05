// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/domain"
	storage "github.com/verbiscms/verbis/api/mocks/services/storage"
	theme "github.com/verbiscms/verbis/api/mocks/services/theme"
	webp "github.com/verbiscms/verbis/api/mocks/services/webp"
	repo "github.com/verbiscms/verbis/api/mocks/store/media"
)

func (t *MediaServiceTestSuite) TestService_ReGenerateWebP() {
	tt := map[string]struct {
		mock func(r *repo.Repository, s *storage.Bucket, th *theme.Service)
		want interface{}
	}{
		"Find Error": {
			func(r *repo.Repository, s *storage.Bucket, th *theme.Service) {
				r.On("List", params.Params{LimitAll: true}).Return(domain.MediaItems{testMedia}, 1, nil)
				s.On("Find", mock.Anything).Return(nil, domain.File{}, fmt.Errorf("error"))
			},
			1,
		},
		"List Error": {
			func(r *repo.Repository, s *storage.Bucket, th *theme.Service) {
				r.On("List", params.Params{LimitAll: true}).Return(nil, 0, fmt.Errorf("error"))
			},
			"error",
		},
		"None Found": {
			func(r *repo.Repository, s *storage.Bucket, th *theme.Service) {
				r.On("List", params.Params{LimitAll: true}).Return(nil, 0, nil)
			},
			"no webp images to process",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(domain.Options{}, test.mock)
			got, err := s.ReGenerateWebP()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestService_GenerateWebP() {
	mockfn := func(r *repo.Repository, s *storage.Bucket, th *theme.Service) {
		s.On("Find", testMedia.File.URL+domain.WebPExtension).Return(nil, domain.File{ID: 1}, nil)
		s.On("Delete", 1).Return(nil)
		s.On("Find", testMedia.File.URL).Return(nil, domain.File{}, nil)
		s.On("Upload", mock.Anything).Return(domain.File{}, nil)
	}

	s := t.Setup(domain.Options{}, mockfn)
	w := &webp.Execer{}
	w.On("Convert", mock.Anything, 0).Return(bytes.NewReader([]byte("test")), nil)

	s.webp = w
	s.generateWebP(domain.MediaItems{testMediaSizes}, 0)
}
