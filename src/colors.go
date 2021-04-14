package main

import (
	"image/color"
)

func getColor(colorName string) color.RGBA {
	switch colorName {
	case "black":
		return color.RGBA{0, 0, 0, 0xff}
	case "white":
		return color.RGBA{255, 255, 255, 0xff}
	case "red":
		return color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	case "green":
		return color.RGBA{0x00, 0xFF, 0x00, 0xFF}
	case "blue":
		return color.RGBA{0x00, 0x00, 0xFF, 0xFF}
	default:
		return color.RGBA{100, 200, 200, 0xff}
	}
}
