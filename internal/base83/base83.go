package base83

import "math"

var digitCharacters = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
	"#",
	"$",
	"%",
	"*",
	"+",
	",",
	"-",
	".",
	":",
	";",
	"=",
	"?",
	"@",
	"[",
	"]",
	"^",
	"_",
	"{",
	"|",
	"}",
	"~",
}

func Encode83(n float64, length int) (result string) {
	for i := 1; i <= length; i++ {
		digit := int(math.Mod(math.Floor(n)/math.Pow(83, float64(length-i)), 83))

		if digit < 0 {
			digit *= -1
		}

		if digit < 0 || digit > len(digitCharacters) {
			digit = 0
		}

		result += digitCharacters[digit]
	}

	return
}

func Decode83(blurHash string) (res float64) {
	blurHashRunes := []rune(blurHash)
	for i := 0; i < len(blurHashRunes); i++ {
		char := blurHashRunes[i]

		var digit int

		for j, c := range digitCharacters {
			if string(char) == c {
				digit = j
			}
		}
		res = res*float64(83) + float64(digit)
	}

	return
}
