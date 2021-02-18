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
