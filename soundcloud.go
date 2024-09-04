package soundcloud

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

type Soundcloud struct {
	trackURL string
	options  *Options
}

type Options struct {
	InvertColors bool
	Watermark    bool
}

func New(trackURL string, options *Options) *Soundcloud {
	return &Soundcloud{
		trackURL: trackURL,
		options:  options,
	}
}

func (sc *Soundcloud) Run() {
	scAPI, err := soundcloudapi.New(soundcloudapi.APIOptions{})
	if err != nil {
		log.Println("soundcloud ERROR:", err)
		return
	}

	if !soundcloudapi.IsURL(sc.trackURL, true, true) {
		log.Println("soundcloud: proper track URL is needed")
		return
	}

	tracks, err := scAPI.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
		URL: sc.trackURL,
	})
	if err != nil {
		log.Println("soundcloud ERROR:", err)
		return
	}
	track := tracks[0]

	log.Println("soundcloud: track:", track.Title)
	log.Println("soundcloud: artist:", track.User.Username)

	name := fmt.Sprintf("%s - %s", track.User.Username, track.Title)

	log.Println("soundcloud: creating remix folder:", name)

	err = os.Mkdir(name, 0777)
	if err != nil {
		log.Println("soundcloud ERROR:", err)
		return
	}

	log.Println("soundcloud: extracting track art")

	trackArt, err := sc.extractTrackArt(track.ArtworkURL)
	if err != nil {
		log.Println("soundcloud ERROR:", err)
		return
	}

	buff := bytes.NewBuffer(trackArt)
	log.Println("soundcloud: extracted track art")

	trackArtImage, err := jpeg.Decode(buff)
	if err != nil {
		log.Println("soundcloud ERROR:", err)
		return
	}

	if sc.options.InvertColors {
		log.Println("soundcloud: inverting colors")
		trackArtImage = invertImageColor(trackArtImage)
		log.Println("soundcloud: inverted colors")
	}

	if sc.options.Watermark {
		log.Println("soundcloud: applying watermark")
		trackArtImage = applyWatermark(trackArtImage)
		log.Println("soundcloud: applied watermark")
	}

	remixTrackArt, err := os.Create(fmt.Sprintf("./%s/trackart.png", name))
	if err != nil {
		log.Println("soundcloud ERROR:", err)
		return
	}
	defer remixTrackArt.Close()

	png.Encode(remixTrackArt, trackArtImage)

	log.Println("soundcloud: saved trackart")

	log.Println("soundcloud: downloading track")

	mp3Out, err := os.Create(fmt.Sprintf("./%s/music.mp3", name))
	if err != nil {
		log.Println("soundcloud ERROR:", err)
		return
	}
	defer mp3Out.Close()

	err = scAPI.DownloadTrack(tracks[0].Media.Transcodings[0], mp3Out)
	if err != nil {
		log.Println("soundcloud ERROR:", err)
		return
	}

	log.Println("soundcloud: saved track")
}
