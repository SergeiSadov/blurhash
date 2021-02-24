package utils

import (
	"testing"
)

func TestGetImg(t *testing.T) {
	type args struct {
		componentX int
		componentY int
		path       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"0", args{1, 8, "../../resources/black.png"}, false},
		{"1", args{1, 8, "../../resources/notexist.png"}, true},
		{"2", args{1, 8, "../../resources/notanimage.png"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetImg(tt.args.componentX, tt.args.componentY, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
