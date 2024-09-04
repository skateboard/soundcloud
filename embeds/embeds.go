package embeds

import (
	"embed"
	"image"
	"image/png"
)

var (
	//go:embed watermark.png
	watermark embed.FS
)

func GetWatermark() (image.Image, error) {
	wM, err := watermark.Open("watermark.png")
	if err != nil {
		return nil, err
	}

	return png.Decode(wM)

}
