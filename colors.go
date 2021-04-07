package main

import (
	"image/color"
)

func getColor(colorName string) color.RGBA {
	if colorName == "black" {
		return color.RGBA{0, 0, 0, 0xff}
	}
	if colorName == "white" {
		return color.RGBA{255, 255, 255, 0xff}
	}
	if colorName == "red" {
		return color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	}
	if colorName == "green" {
		return color.RGBA{0x00, 0xFF, 0x00, 0xFF}
	}
	if colorName == "blue" {
		return color.RGBA{0x00, 0x00, 0xFF, 0xFF}
	}
	return color.RGBA{100, 200, 200, 0xff}
}
