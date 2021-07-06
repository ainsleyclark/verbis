// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resizer

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/services/media/internal/image"
	"github.com/ainsleyclark/verbis/api/storage"
	"github.com/disintegration/imaging"
	stdimage "image"
)

// Resizer describes the method for resizing images for
// the library.
type Resizer interface {
	Resize(imager image.Imager, dest string, media domain.MediaSize) (domain.File, error)
}

// Resize defines the data needed for resizing images.
type Resize struct {
	Storage     storage.Client
	Compression int
}

// Resize satisfies the Resizer by decoding, cropping and
// resizing and finally saving the resized image.
func (r *Resize) Resize(imager image.Imager, dest string, media domain.MediaSize) (domain.File, error) {
	i, err := imager.Decode()
	if err != nil {
		return domain.File{}, err
	}

	var resized *stdimage.NRGBA
	if media.Crop {
		resized = imaging.Fill(i, media.Width, media.Height, imaging.Center, imaging.Lanczos)
	} else {
		resized = imaging.Resize(i, media.Width, media.Height, imaging.Lanczos)
	}

	enc, err := imager.Encode(resized, r.Compression)
	if err != nil {
		return domain.File{}, err
	}

	upload, err := r.Storage.Upload(dest, int64(enc.Len()), enc)
	if err != nil {
		return domain.File{}, err
	}

	return upload, nil
}
