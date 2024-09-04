package soundcloud

import (
	"image"
	"image/color"
)

func invertImageColor(srcImage image.Image) image.Image {

	// Create a new image
	bounds := srcImage.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	rectangle := image.Rect(0, 0, w, h)
	newImage := image.NewRGBA(rectangle)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {

			oldColor := srcImage.At(x, y)
			red, green, blue, alpha := oldColor.RGBA()

			// transforming pixel into inverse.
			ired := 255 - (uint8(red >> 8))
			igreen := 255 - (uint8(green >> 8))
			iblue := 255 - (uint8(blue >> 8))
			newColor := color.RGBA{ired, igreen, iblue, uint8(alpha >> 8)}

			newImage.Set(x, y, newColor)
		}
	}

	return newImage
}
