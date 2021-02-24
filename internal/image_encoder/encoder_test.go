package image_encoder

import (
	"bytes"
	"image"
	"testing"
)

func TestEncodeImg(t *testing.T) {
	type args struct {
		ext string
		img image.Image
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"wrong ext", args{".test", image.NewNRGBA(image.Rect(0, 0, 0, 0))}, true},
		{"png", args{pngExt, image.NewNRGBA(image.Rect(0, 0, 10, 10))}, false},
		{"jpg", args{jpgExt, image.NewNRGBA(image.Rect(0, 0, 0, 0))}, false},
		{"jpeg", args{jpegExt, image.NewNRGBA(image.Rect(0, 0, 0, 0))}, false},
		{"gif", args{gifExt, image.NewNRGBA(image.Rect(0, 0, 0, 0))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := EncodeImg(tt.args.ext, w, tt.args.img)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeImg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
