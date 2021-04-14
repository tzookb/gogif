package main

// https://github.com/urfave/cli/blob/master/docs/v2/manual.md#getting-started

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/uniplaces/carbon"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

func getImages() []image.Image {
	then, _ := getThen()
	now, _ := carbon.NowInLocation("Local")
	size := 2
	timeBlocks := getDiffsBack(size, now, then)

	allImages := make([]image.Image, size)

	for i, tb := range *timeBlocks {
		curImg := createImageFrame()
		days := tb.days
		hours := tb.hours
		minutes := tb.minutes
		seconds := tb.seconds
		timeString := fmt.Sprintf("heyy %d:%d:%d:%d", days, hours, minutes, seconds)
		// imgFromFile, _ := getImageFromFilePath("./src/images/template.jpg")
		drawText(curImg, timeString)
		// fmt.Printf("type %T\n", curImg)
		allImages[i] = curImg
	}

	return allImages
}

func createImageFrame() *image.RGBA {
	width := 600
	height := 600

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	imgframe := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	draw.Draw(imgframe, image.Rect(0, 0, width, height), &image.Uniform{getColor("blue")}, image.ZP, draw.Src)

	return imgframe
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)

	return image, err
}

func CreateBasicGif(out io.Writer, palette []image.Paletted) {
	delays := make([]int, len(palette))
	pointerPallets := make([]*image.Paletted, len(palette))
	for i := range palette {
		delays[i] = 100
		pointerPallets[i] = &palette[i]
	}
	anim := gif.GIF{Delay: delays, Image: pointerPallets}

	gif.EncodeAll(out, &anim)
}

type ImgProcResp struct {
	position int
	img      image.Paletted
}

func EncodeImgPaletted(images []image.Image) []image.Paletted {
	// Gif options
	opt := gif.Options{}
	g := make([]image.Paletted, len(images))
	channel := make(chan ImgProcResp)

	for pos, im := range images {
		go func(pos int, im image.Image) {
			b := bytes.Buffer{}
			// Write img2gif file to buffer.
			err := gif.Encode(&b, im, &opt)

			if err != nil {
				println(err)
			}
			// Decode img2gif file from buffer to img.
			img, err := gif.Decode(&b)

			if err != nil {
				println(err)
			}

			// Cast img.
			theItem, _ := img.(*image.Paletted)

			channel <- ImgProcResp{pos, *theItem}
		}(pos, im)
	}

	i := 0
	for i < len(images) {
		theItem := <-channel
		g[theItem.position] = theItem.img
		// g = append(g, theItem)
		i += 1
	}
	return g
}

func drawText(canvas *image.RGBA, text string) error {
	fontSize := 25.0
	// fgColor = image.White
	fgColor := image.NewUniform(getColor("red"))

	fontFace, err := freetype.ParseFont(goregular.TTF)
	fontDrawer := &font.Drawer{
		Dst: canvas,
		Src: fgColor,
		Face: truetype.NewFace(fontFace, &truetype.Options{
			Size:    fontSize,
			Hinting: font.HintingFull,
		}),
	}
	textBounds, _ := fontDrawer.BoundString(text)
	xPosition := (fixed.I(canvas.Rect.Max.X) - fontDrawer.MeasureString(text)) / 2
	textHeight := textBounds.Max.Y - textBounds.Min.Y
	yPosition := fixed.I((canvas.Rect.Max.Y)-textHeight.Ceil())/2 + fixed.I(textHeight.Ceil())
	fontDrawer.Dot = fixed.Point26_6{
		X: xPosition,
		Y: yPosition,
	}
	fontDrawer.DrawString(text)
	return err
}
