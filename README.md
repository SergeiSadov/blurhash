# blurhash [![Go Report Card](https://goreportcard.com/badge/github.com/sergeisadov/blurhash)](https://goreportcard.com/report/github.com/sergeisadov/blurhash) [![license](https://img.shields.io/github/license/sergeisadov/blurhash.svg)](https://github.com/sergeisadov/blurhash/blob/master/LICENSE) [![Build Status](https://travis-ci.org/sergeisadov/blurhash.svg)](https://travis-ci.org/sergeisadov/blurhash) [![codecov](https://codecov.io/gh/SergeiSadov/blurhash/branch/main/graph/badge.svg)](https://codecov.io/gh/SergeiSadov/blurhash)

# BlurHash encoder and decoder in GO

This is basically a port of the [Wolt TypeScript](https://github.com/woltapp/blurhash/tree/master/TypeScript) version of
the [Blurhash alorithm](https://github.com/woltapp/blurhash/blob/master/Algorithm.md) implementation

At the moment works with png only

## Usage as a library

Just import package using go get:

`go get -u github.com/sergeysadov/blurhash`

For encoding use:

```go
img, err := utils.GetImg(componentX, componentY, path)
if err != nil {
//handle error
}

hash, _ := blurhash.Encode(img)
```

For decoding use:

```go
pixels, err := blurhash.Decode(blurHash, width, height, 0)
if err != nil {
//handle error
}

nrgba := image.NewRGBA(image.Rect(0, 0, width, height))
nrgba.Pix = pixels

out, err := os.Create(resultPath)
if err != nil {
//handler error
}

defer out.Close()

if err := png.Encode(out, nrgba); err != nil {
//handle error
}
```

## Usage as a tool

To build encoder use

```shell
	 $ make encoder
	 $ ./blurhash-encoder -xComponents 4 -yComponents 3 -path resources/pic.png
	 LaJR8MVu8_~po#smR+a~xaoLWCRj
```

To build decoder use

```shell
	 $ make decoder
	 $ ./blurhash-decoder -width 32 -height 32 -hash LaJR8MVu8_~po#smR+a~xaoLWCRj -result result.png
	 LaJR8MVu8_~po#smR+a~xaoLWCRj
```

### TODO List

- [ ] Other image formats
- [ ] Benchmarks
- [ ] More unit tests

## Licence

This project is licensed under the [MIT License](LICENSE)
