// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	resizer "github.com/verbiscms/verbis/api/mocks/services/media/resizer"
	storage "github.com/verbiscms/verbis/api/mocks/services/storage"
	repo "github.com/verbiscms/verbis/api/mocks/store/media"
	"github.com/verbiscms/verbis/api/test"
	"github.com/verbiscms/verbis/api/test/dummy"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// MediaServiceTestSuite defines the helper used for testMedia
// library testing.
type MediaServiceTestSuite struct {
	test.MediaSuite
	TestDataPath string
	LogWriter    *bytes.Buffer
}

// TestMediaService asserts testing has begun.
func TestMediaService(t *testing.T) {
	suite.Run(t, &MediaServiceTestSuite{
		MediaSuite: test.NewMediaSuite(),
	})
}

// BeforeTest setups the LogWriter
func (t *MediaServiceTestSuite) BeforeTest(suiteName, testName string) {
	b := &bytes.Buffer{}
	t.LogWriter = b
	logger.SetOutput(b)
}

// SetupSuite reassigns API path for testing.
func (t *MediaServiceTestSuite) SetupSuite() {
	wd, err := os.Getwd()
	t.NoError(err)
	t.TestDataPath = filepath.Join(wd, "testdata")
}

// Reset the log writer.
func (t *MediaServiceTestSuite) Reset() {
	t.LogWriter.Reset()
}

// File returns a byte value of a path in the testdata
// directory for testing.
func (t *MediaServiceTestSuite) File(path string) []byte {
	b, err := ioutil.ReadFile(filepath.Join(t.TestDataPath, path))
	t.NoError(err)
	return b
}

const (
	// MediaID is the default ID use for testing.
	MediaID = 1
	// TestFileURL defines the URL for media items for testing.
	TestFileURL = "/uploads/2020/01/file.jpg"
	// TestFileURL defines the WebP URL for media items for
	// testing.
	TestFileURLWebP = TestFileURL + domain.WebPExtension
)

var (
	// svgFile is the default domain.File with an SVG mime
	// type used for testing.
	svgFile = domain.File{
		Id:       1,
		Url:      "/uploads/gopher.svg",
		Name:     "gopher.svg",
		BucketId: "/uploads/gopher.svg",
		Mime:     "image/svg+xml",
		Private:  false,
	}
	// pngFile is the default domain.File with an PNG mime
	// type used for testing.
	pngFile = domain.File{
		Id:       1,
		Url:      "/uploads/gopher.png",
		Name:     "gopher.png",
		BucketId: "/uploads/gopher.png",
		Mime:     "image/png",
		Private:  false,
	}
	// jpgFile is the default domain.File with an JPG mime
	// type used for testing.
	jpgFile = domain.File{
		Id:       1,
		Url:      "/uploads/gopher.jpg",
		Name:     "gopher.jpg",
		BucketId: "/uploads/gopher.jpg",
		Mime:     "image/jpeg",
		Private:  false,
	}
	// testMedia is the default media Item used for
	// testing.
	testMedia = domain.Media{
		Id:   MediaID,
		File: domain.File{Id: 1, Url: TestFileURL},
	}
	// testMediaSlice are the default media items used
	// for testing.
	testMediaSlice = domain.MediaItems{
		domain.Media{
			Id:   MediaID,
			File: domain.File{Id: 1, Url: TestFileURL},
		},
		domain.Media{
			Id:   MediaID,
			File: domain.File{Id: 1, Url: TestFileURL},
		},
	}
	// testMediaSizes are the default media sizes used
	// for testing.
	testMediaSizes = domain.Media{
		Id: MediaID,
		Sizes: domain.MediaSizes{
			"thumnbnail": domain.MediaSize{
				SizeKey:  "key",
				SizeName: "name",
				File:     domain.File{Id: 1, Url: TestFileURL},
			},
		},
		File: domain.File{Id: 1, Url: TestFileURL},
	}
	// opts is the default options with media sizes used
	// for testing.
	opts = &domain.Options{
		MediaSizes: domain.MediaSizes{"thumbnail": domain.MediaSize{SizeKey: "thumb", SizeName: "thumb", Width: 300, Height: 300, Crop: false}},
	}
)

// Setup is a helper to obtain a mock testMedia Service
// for testing.
func (t *MediaServiceTestSuite) Setup(cfg *domain.ThemeConfig, opts *domain.Options, mock func(r *repo.Repository, s *storage.Bucket)) *Service {
	r := &repo.Repository{}
	s := &storage.Bucket{}

	if mock != nil {
		mock(r, s)
	}

	if cfg == nil {
		cfg = &domain.ThemeConfig{}
	}

	if opts == nil {
		opts = &domain.Options{}
	}

	c := New(opts, s, r)
	c.resizer = &resizer.Resizer{}
	c.config = cfg

	return c
}

func (t *MediaServiceTestSuite) TestService_List() {
	tt := map[string]struct {
		mock func(r *repo.Repository, s *storage.Bucket)
		want interface{}
	}{
		"Success": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("List", dummy.DefaultParams).Return(testMediaSlice, 2, nil)
			},
			testMediaSlice,
		},
		"Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("List", dummy.DefaultParams).Return(nil, 0, fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, nil, test.mock)
			got, _, err := s.List(dummy.DefaultParams)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestService_Find() {
	tt := map[string]struct {
		mock func(r *repo.Repository, s *storage.Bucket)
		want interface{}
	}{
		"Success": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Find", 1).Return(testMedia, nil)
			},
			testMedia,
		},
		"Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Find", 1).Return(domain.Media{}, fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, nil, test.mock)
			got, err := s.Find(1)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestService_Update() {
	tt := map[string]struct {
		mock func(r *repo.Repository, s *storage.Bucket)
		want interface{}
	}{
		"Success": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Update", testMedia).Return(testMedia, nil)
			},
			testMedia,
		},
		"Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Update", testMedia).Return(domain.Media{}, fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, nil, test.mock)
			got, err := s.Update(testMedia)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
