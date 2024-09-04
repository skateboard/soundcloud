package soundcloud

import (
	"image"
	"image/draw"

	"github.com/skateboard/soundcloud/embeds"
)

func applyWatermark(inputImage image.Image) image.Image {
	watermarkImage, err := embeds.GetWatermark()
	if err != nil {
		return nil
	}

	offset := image.Pt(0, 0)
	b := inputImage.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, inputImage, image.ZP, draw.Src)
	draw.Draw(m, watermarkImage.Bounds().Add(offset), watermarkImage, image.ZP, draw.Over)

	return m
}
