// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/verbiscms/verbis/api/common/mime"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"image"
	"mime/multipart"
)

//
//
// Validate satisfies the Library to see if the media item passed
// and is valid for upload.
func (s *Service) Validate(file *multipart.FileHeader) error {
	const op = "Media.Validate"

	cfg, err := s.theme.Config()
	if err != nil {
		return err
	}

	mt, err := file.Open()
	defer func() {
		err := mt.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing file with the name: " + file.Filename, Operation: op, Err: err})
		}
	}()

	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error opening file with the name: " + file.Filename, Operation: op, Err: err}
	}

	v := validator{
		Config:  cfg,
		Options: s.options.Struct(),
		Size:    file.Size,
		File:    mt,
	}

	err = v.Mime()
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "The file is not permitted to be uploaded", Operation: op, Err: err}
	}

	err = v.FileSize()
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "The file exceeds the maximum size restriction", Operation: op, Err: err}
	}

	err = v.Image()
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "The image exceeds the width/height restrictions", Operation: op, Err: err}
	}

	return nil
}

// validator defines the helper for validating media items.
type validator struct {
	Config  domain.ThemeConfig
	Options domain.Options
	Size    int64
	File    multipart.File
}

//// validate Checks for valid mime types, file sizes and
//// image sizes
//// Returns errors.INVALID if any condition is not met.
//func validate(h *multipart.FileHeader, opts domain.Options, cfg domain.ThemeConfig) error {
//
//	return nil
//}

// Mime checks if a mimetype is valid by comparing the
// mime with the allowed file types in the configuration.
// Returns ErrMimeType on failure.
func (v *validator) Mime() error {
	m, err := mimetype.DetectReader(v.File)
	if err != nil {
		return ErrMimeType
	}

	valid := mime.IsValidMime(v.Config.Media.AllowedFileTypes, m.String())
	if !valid {
		return ErrMimeType
	}

	return nil
}

// FileSize checks if the file size is under the upload
// maximum size set in the options.
// Returns ErrFileTooBig on failure.
func (v *validator) FileSize() error {
	fileSize := v.Size / 1024
	if fileSize > v.Options.MediaUploadMaxSize && v.Options.MediaUploadMaxSize != 0 {
		return ErrFileTooBig
	}
	return nil
}

// Image checks if an image is over the maximum width and
// height set in the options. Returns nil if the file is
// not an image.
func (v *validator) Image() error {
	_, err := v.File.Seek(0, 0)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(v.File)
	if err != nil {
		return nil // Is not an image
	}

	bounds := img.Bounds().Max
	if int64(bounds.X) > v.Options.MediaUploadMaxWidth && v.Options.MediaUploadMaxWidth != 0 {
		return errors.New("image exceeds the maximum upload width")
	}

	if int64(bounds.Y) > v.Options.MediaUploadMaxHeight && v.Options.MediaUploadMaxHeight != 0 {
		return errors.New("image exceeds the maximum upload height")
	}

	return nil
}
