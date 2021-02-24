package blurhash

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		blurHash string
		width    int
		height   int
		punch    int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"0", args{"LRH2DzI9-;M_~pI9jEjY-pNabatR", 32, 32, 0}, false},
		{"1", args{"L009jvfQfQfQfQfQfQfQfQfQfQfQ", 32, 32, 1}, false},
		{"2", args{"L2TSUA~qfQ~q~qj[fQj[fQfQfQfQ", 32, 32, 0}, false},
		{"3", args{"TRFSE~qfQ~q~qj[fQj[fQfQfQfQ", 32, 32, 0}, true},
		{"4", args{"", 32, 32, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Decode(tt.args.blurHash, tt.args.width, tt.args.height, tt.args.punch)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_validateBlurhash(t *testing.T) {
	type args struct {
		blurhash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"0", args{""}, true},
		{"1", args{"LRH2DzI9-;M_~pI9jEjY-pNabatR"}, false},
		{"2", args{"LRH2Ed#5:5"}, true},
		{"3", args{"L009jvfQfQfQfQfQfQfQfQfQfQfQ"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateBlurhash(tt.args.blurhash); (err != nil) != tt.wantErr {
				t.Errorf("validateBlurhash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isBlurhashValid(t *testing.T) {
	type args struct {
		blurHash string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{""}, false},
		{"1", args{"LRH2DzI9-;M_~pI9jEjY-pNabatR"}, true},
		{"2", args{"LRH2Ed#5:5"}, false},
		{"3", args{"L009jvfQfQfQfQfQfQfQfQfQfQfQ"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBlurhashValid(tt.args.blurHash); got != tt.want {
				t.Errorf("isBlurhashValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeDC(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"0", args{-1}, []float64{-0.0003035269835488375, 1, 1}},
		{"1", args{0}, []float64{0, 0, 0}},
		{"2", args{1}, []float64{0, 0, 0.0003035269835488375}},
		{"3", args{255}, []float64{0, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeDC(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeDC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeAC(t *testing.T) {
	type args struct {
		value        float64
		maximumValue float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"0", args{-1, 0.1}, []float64{-0.12345679012345681, -0.12345679012345681, -0.12345679012345681}},
		{"1", args{0, 0.1}, []float64{-0.1, -0.1, -0.1}},
		{"2", args{1, 0.1}, []float64{-0.1, -0.1, -0.07901234567901234}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeAC(tt.args.value, tt.args.maximumValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeAC() = %v, want %v", got, tt.want)
			}
		})
	}
}
