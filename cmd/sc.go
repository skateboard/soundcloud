package main

import (
	"flag"
	"log"

	"github.com/skateboard/soundcloud"
)

func main() {
	trackUrl := flag.String("track", "track url", "the track url")
	invertColors := flag.Bool("invert", true, "invert the track art colors")
	watermark := flag.Bool("watermark", true, "add watermark to the track art")

	flag.Parse()

	if trackUrl == nil {
		log.Println("soundcloud: track url is needed")
		return
	}

	if invertColors == nil {
		log.Println("soundcloud: invert colors option is needed")
		return
	}

	if watermark == nil {
		log.Println("soundcloud: watermark option is needed")
		return
	}

	sc := soundcloud.New(*trackUrl, &soundcloud.Options{
		InvertColors: *invertColors,
		Watermark:    *watermark,
	})
	sc.Run()
}
