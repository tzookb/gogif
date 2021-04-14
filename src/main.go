package main

// https://github.com/urfave/cli/blob/master/docs/v2/manual.md#getting-started

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	allImages := getImages()
	fmt.Println(len(allImages))
	fmt.Println(len(allImages))
	fmt.Println(len(allImages))
	pallets := EncodeImgPaletted(allImages)

	f, _ := os.Create("image.gif")
	defer f.Close()
	// fmt.Println(pallets)
	CreateBasicGif(f, pallets)
}
