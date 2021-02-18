package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sergeisadov/blurhash/pkg/blurhash"
	"github.com/sergeisadov/blurhash/pkg/utils"
)

func main() {
	var (
		xComponents int
		yComponents int
		path        string
	)

	flag.IntVar(&xComponents, "xComponents", 4, "Number of X components. Must be in interval from 1 to 9")
	flag.IntVar(&yComponents, "yComponents", 3, "Number of Y components. Must be in interval from 1 to 9")
	flag.StringVar(&path, "path", "resources/pic.png", "path")
	flag.Parse()

	hash, err := blurHashForFile(xComponents, yComponents, path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)
}

func blurHashForFile(componentX, componentY int, path string) (hash string, err error) {
	img, err := utils.GetImg(componentX, componentY, path)
	if err != nil {
		return
	}

	return blurhash.Encode(img)
}
