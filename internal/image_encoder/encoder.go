package image_encoder

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

const (
	pngExt  = ".png"
	jpgExt  = ".jpg"
	jpegExt = ".jpeg"
	gifExt  = ".gif"
)

// ErrUnknownImgType unknown image type
var ErrUnknownImgType = errors.New("unknown image type")

// EncodeImg is used to encode img to desirable extension
func EncodeImg(ext string, w io.Writer, img image.Image) error {
	switch ext {
	case pngExt:
		return png.Encode(w, img)
	case jpgExt, jpegExt:
		return jpeg.Encode(w, img, nil)
	case gifExt:
		return gif.Encode(w, img, nil)
	}

	return ErrUnknownImgType
}
