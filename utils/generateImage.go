package utils

import (
	"image"
	"image/color"
)

func GenerateImage(matrix [][]bool) image.Image {
	size := len(matrix)
	scale := 10
	img := image.NewRGBA(image.Rect(0, 0, size*scale, size*scale))
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	for y := range size {
		for x := range size {
			c := white
			if matrix[y][x] {
				c = black
			}
			for dy := range scale {
				for dx := range scale {
					img.Set(x*scale+dx, y*scale+dy, c)
				}
			}
		}
	}
	return img
}
