package soundcloud

import (
	"errors"
	"io"
	"net/http"

	"github.com/johnreutersward/opengraph"
)

func (s *Soundcloud) extractTrackArt() ([]byte, error) {
	metaData, err := fetchMetaData(s.trackURL)
	if err != nil {
		return nil, err
	}

	artwork := extractMetaData(metaData)

	if artwork.ImageURL != "" {
		res, err := http.Get(artwork.ImageURL)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return b, nil
	}

	return nil, errors.New("failed to extract track art")
}

type artwork struct {
	Title    string
	ImageURL string
}

func fetchMetaData(url string) ([]opengraph.MetaData, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	md, err := opengraph.Extract(res.Body)
	if err != nil {
		return nil, err
	}

	return md, nil
}

func extractMetaData(md []opengraph.MetaData) artwork {
	var artwork artwork

	for i := range md {
		switch md[i].Property {
		case "title":
			artwork.Title = md[i].Content
		case "image":
			artwork.ImageURL = md[i].Content
		}
	}

	return artwork
}
