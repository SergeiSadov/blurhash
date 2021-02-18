package blurhash

import (
	"math"
	"strings"

	"github.com/sergeisadov/blurhash/internal/base83"
	"github.com/sergeisadov/blurhash/pkg/utils"
)

const bytesPerPixel = 4

const (
	minComponent = 1
	maxComponent = 9
)

func Encode(img *utils.Img) (string, error) {
	if img.ComponentX < minComponent || img.ComponentX > maxComponent || img.ComponentY < minComponent || img.ComponentY > maxComponent {
		return "", ErrIncorrectComponents
	}

	if img.Width*img.Height*4 != len(img.Pixels) {
		return "", ErrWrongSize
	}

	factors := make([][]float64, 0, img.ComponentX)
	for i := range factors {
		factors[i] = make([]float64, 0, 3)
	}

	hashSB := strings.Builder{}

	for y := 0; y < img.ComponentY; y++ {
		for x := 0; x < img.ComponentX; x++ {
			normalization := 2
			if x == 0 && y == 0 {
				normalization = 1
			}

			factor := multiplyBasisFunction(img.Pixels, img.Width, img.Height, func(i, j int) float64 {
				return float64(normalization) *
					math.Cos(math.Pi*float64(x)*float64(i)/float64(img.Width)) *
					math.Cos(math.Pi*float64(y)*float64(j)/float64(img.Height))

			})

			factors = append(factors, factor)
		}

	}

	dc := factors[0]
	ac := factors[1:]

	hashSB.Grow(1 + 1 + 4 + 2*(len(factors)-1)*9)

	sizeFlag := img.ComponentX - 1 + (img.ComponentY-1)*9

	hashSB.WriteString(base83.Encode83(float64(sizeFlag), 1))

	var maximumValue float64
	if len(ac) > 0 {
		actualMaximumValue := max(ac)
		quantisedMaximumValue := math.Floor(math.Max(0, math.Min(82, math.Floor(actualMaximumValue*166-0.5))))

		maximumValue = (quantisedMaximumValue + 1) / 166
		hashSB.WriteString(base83.Encode83(math.Round(quantisedMaximumValue), 1))
	} else {
		maximumValue = 1
		hashSB.WriteString(base83.Encode83(0, 1))
	}

	hashSB.WriteString(base83.Encode83(float64(encodeDC(dc)), 4))

	for i := range ac {
		hashSB.WriteString(base83.Encode83(encodeAC(ac[i], maximumValue), 2))
	}

	return hashSB.String(), nil
}

func sRGBToLinear(value float64) float64 {
	v := value / 255
	if v <= 0.04045 {
		return v / 12.92
	}

	return math.Pow((v+0.055)/1.055, 2.4)
}

func linearTosRGB(value float64) float64 {
	v := math.Max(0, math.Min(1, value))

	if v <= 0.0031308 {
		return math.Round(v*12.92*255 + 0.5)
	}

	return math.Round((1.055*math.Pow(v, 1/2.4)-0.055)*255 + 0.5)
}

func sign(n float64) float64 {
	if n < 0 {
		return -1
	}

	return 1
}

func signPow(val, exp float64) float64 {
	return sign(val) * math.Pow(math.Abs(val), exp)
}

func multiplyBasisFunction(pixels []uint8, width, height int, basisFunc func(i, j int) float64) []float64 {
	r := float64(0)
	g := float64(0)
	b := float64(0)

	bytesPerRow := width * bytesPerPixel

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			basis := basisFunc(x, y)
			r += basis * sRGBToLinear(float64(pixels[bytesPerPixel*x+0+y*bytesPerRow]))
			g += basis * sRGBToLinear(float64(pixels[bytesPerPixel*x+1+y*bytesPerRow]))
			b += basis * sRGBToLinear(float64(pixels[bytesPerPixel*x+2+y*bytesPerRow]))
		}
	}

	scale := float64(1) / (float64(width) * float64(height))

	return []float64{r * scale, g * scale, b * scale}
}

func max(input [][]float64) float64 {
	var max float64
	for i := range input {
		for j := range input[i] {
			if max < input[i][j] {
				max = input[i][j]
			}

		}
	}
	return max
}

func encodeDC(val []float64) int {
	roundedR := int(linearTosRGB(val[0]))
	roundedG := int(linearTosRGB(val[1]))
	roundedB := int(linearTosRGB(val[2]))

	return (roundedR << 16) + (roundedG << 8) + roundedB
}

func encodeAC(value []float64, maximumValue float64) float64 {
	exponent := 0.5
	minX := float64(18)

	quantR := math.Floor(math.Max(0,
		math.Min(minX, math.Floor(signPow(value[0]/maximumValue, exponent)*9+9.5))))

	quantG := math.Floor(math.Max(0,
		math.Min(minX, math.Floor(signPow(value[1]/maximumValue, exponent)*9+9.5))))

	quantB := math.Floor(math.Max(0,
		math.Min(minX, math.Floor(signPow(value[2]/maximumValue, exponent)*9+9.5))))

	return quantR*19*19 + quantG*19 + quantB
}
