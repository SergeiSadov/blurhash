package base83

import (
	"testing"
)

func Test_Encode83(t *testing.T) {
	type args struct {
		n      float64
		length int
	}

	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{"encode to L", args{21, 1}, "L"},
		{"encode to 0000", args{0, 4}, "0000"},
		{"encode to 0001", args{1, 4}, "0001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Encode83(tt.args.n, tt.args.length); gotResult != tt.wantResult {
				t.Errorf("Encode83() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestDecode83(t *testing.T) {
	type args struct {
		blurHash string
	}
	tests := []struct {
		name    string
		args    args
		wantRes float64
	}{
		{"0", args{""}, 0},
		{"1", args{"L"}, 21},
		{"2", args{"0000"}, 0},
		{"3", args{"0001"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Decode83(tt.args.blurHash); gotRes != tt.wantRes {
				t.Errorf("Decode83() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
