package blurhash

import (
	"fmt"
	"math"

	"github.com/sergeisadov/blurhash/internal/base83"
)

// Decode method takes blurhash string, with, height and punch as params and returs pixels byte slice
// blurhash must be valid to pass method's validation
func Decode(blurHash string, width, height, punch int) (pixels []uint8, err error) {
	if !isBlurhashValid(blurHash) {
		return pixels, ErrBlurhashInvalid
	}

	if punch == 0 {
		punch = 1
	}

	blurHashRunes := []rune(blurHash)

	sizeFlag := base83.Decode83(string(blurHashRunes[0]))
	numY := int(math.Floor(sizeFlag/9) + 1)
	numX := int((math.Mod(sizeFlag, 9)) + 1)

	quantisedMaximumValue := base83.Decode83(string(blurHashRunes[1]))
	maximumValue := (quantisedMaximumValue + 1) / 166

	colorsLen := numX * numY
	colors := make([][]float64, colorsLen)

	for i := 0; i < colorsLen; i++ {
		if i == 0 {
			value := base83.Decode83(string(blurHashRunes[2:6]))
			colors[i] = decodeDC(int(value))
		} else {
			value := base83.Decode83(string(blurHashRunes[4+i*2 : 6+i*2]))
			colors[i] = decodeAC(value, maximumValue*float64(punch))
		}
	}

	bytesPerRow := width * 4
	pixels = make([]uint8, bytesPerRow*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b := float64(0), float64(0), float64(0)

			for j := 0; j < numY; j++ {
				for i := 0; i < numX; i++ {
					basis := math.Cos((math.Pi*float64(x)*float64(i))/float64(width)) *
						math.Cos((math.Pi*float64(y)*float64(j))/float64(height))

					color := colors[i+j*numX]

					r += color[0] * basis
					g += color[1] * basis
					b += color[2] * basis

				}
			}

			intR := linearTosRGB(r)
			intG := linearTosRGB(g)
			intB := linearTosRGB(b)

			pixels[4*x+0+y*bytesPerRow] = uint8(intR)
			pixels[4*x+1+y*bytesPerRow] = uint8(intG)
			pixels[4*x+2+y*bytesPerRow] = uint8(intB)
			pixels[4*x+3+y*bytesPerRow] = 255

		}
	}

	return
}

func validateBlurhash(blurhash string) error {
	blurhashRunes := []rune(blurhash)

	blurhashLen := len(blurhashRunes)
	if blurhash == "" || blurhashLen < 6 {
		return ErrWrongBlurhashLen
	}

	sizeFlag := base83.Decode83(string(blurhashRunes[0]))
	numY := math.Floor(sizeFlag/9) + 1
	numX := (math.Mod(sizeFlag, 9)) + float64(1)

	expectedLen := 4 + 2*numX*numY
	if blurhashLen != int(expectedLen) {
		return fmt.Errorf("blurhash length mismatch: length is %d, but it should be %f", blurhashLen, expectedLen)
	}

	return nil
}

func isBlurhashValid(blurHash string) bool {
	if err := validateBlurhash(blurHash); err != nil {
		return false
	}

	return true
}

func decodeDC(value int) []float64 {
	intR := value >> 16
	intG := (value >> 8) & 255
	intB := value & 255
	return []float64{sRGBToLinear(float64(intR)), sRGBToLinear(float64(intG)), sRGBToLinear(float64(intB))}
}

func decodeAC(value float64, maximumValue float64) []float64 {
	quantR := math.Floor(value / (float64(19) * float64(19)))
	quantG := math.Mod(math.Floor(value/float64(19)), 19)
	quantB := math.Mod(value, 19)

	rgb := []float64{
		signPow((quantR-9)/9, 2.0) * maximumValue,
		signPow((quantG-9)/9, 2.0) * maximumValue,
		signPow((quantB-9)/9, 2.0) * maximumValue,
	}

	return rgb
}
