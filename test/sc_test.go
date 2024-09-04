package test_test

import (
	"testing"

	"github.com/skateboard/soundcloud"
)

func TestSC(t *testing.T) {
	sc := soundcloud.New("https://soundcloud.com/ilovelilshine/trending-topic-prod-solxmn", &soundcloud.Options{
		InvertColors: true,
		Watermark:    true,
	})
	sc.Run()
}
