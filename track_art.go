package soundcloud

import (
	"io"
	"net/http"
	"strings"
)

func (s *Soundcloud) extractTrackArt(artWorkUrl string) ([]byte, error) {
	artWorkUrl = strings.ReplaceAll(artWorkUrl, "large", "t500x500")

	res, err := http.Get(artWorkUrl)
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
