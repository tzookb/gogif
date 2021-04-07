package main

// https://github.com/urfave/cli/blob/master/docs/v2/manual.md#getting-started

import (
	"os"
)

func main() {
	allImages := getImages()
	pallets := EncodeImgPaletted(allImages)

	f, _ := os.Create("image.gif")
	defer f.Close()
	CreateBasicGif(f, pallets)
}
