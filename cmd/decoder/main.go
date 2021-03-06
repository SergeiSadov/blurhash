package main

import (
	"flag"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path"

	"github.com/sergeisadov/blurhash/internal/image_encoder"
	"github.com/sergeisadov/blurhash/pkg/blurhash"
)

func main() {
	var (
		width      int
		height     int
		blurHash   string
		resultPath string
	)

	flag.IntVar(&width, "width", 32, "width")
	flag.IntVar(&height, "height", 32, "height")
	flag.StringVar(&blurHash, "hash", "LRHLMAI9-;Mw~pI9nhjY-pNabtt8", "hash")
	flag.StringVar(&resultPath, "result", "result.png", "result image name")

	flag.Parse()

	pixels, err := blurhash.Decode(blurHash, width, height, 0)
	if err != nil {
		log.Fatal(err)
	}

	nrgba := image.NewRGBA(image.Rect(0, 0, width, height))
	nrgba.Pix = pixels

	out, err := os.Create(resultPath)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	if err := image_encoder.EncodeImg(path.Ext(resultPath), out, nrgba); err != nil {
		log.Fatalf("encoding err: %v", err)
	}
}
