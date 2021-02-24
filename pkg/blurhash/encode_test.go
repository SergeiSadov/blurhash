package blurhash

import (
	"testing"

	"github.com/sergeisadov/blurhash/pkg/utils"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name    string
		imgPath string
		want    string
	}{
		{
			name:    "0",
			imgPath: "../../resources/pic.png",
			want:    "LRHBxaI9-;Mx~pI9nhjY-pNabatR",
		},
		{
			name:    "1",
			imgPath: "../../resources/black.png",
			want:    "L009jvfQfQfQfQfQfQfQfQfQfQfQ",
		},
		{
			name:    "2",
			imgPath: "../../resources/white.png",
			want:    "LETSUA_3fQ_3~qoffQoffQfQfQfQ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := utils.GetImg(4, 3, tt.imgPath)
			if err != nil {
				t.Fatal(err)
			}

			got, err := Encode(img)
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sRGBToLinear(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"0", args{0}, 0},
		{"1", args{0.1}, 0.00003035269835488375},
		{"2", args{1}, 0.0003035269835488375},
		{"3", args{255}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sRGBToLinear(tt.args.value); got != tt.want {
				t.Errorf("sRGBToLinear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linearTosRGB(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"0", args{0}, 1},
		{"1", args{0.1}, 90},
		{"2", args{1}, 255},
		{"3", args{255}, 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linearTosRGB(tt.args.value); got != tt.want {
				t.Errorf("linearTosRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sign(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"0", args{-1}, -1},
		{"1", args{0}, 1},
		{"2", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sign(tt.args.n); got != tt.want {
				t.Errorf("sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
