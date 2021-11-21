package image

import (
	"image"
	"image/png"

	"net/http"
	"os"
	"path/filepath"

	"github.com/canhlinh/svg2png"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func ConvertSvgToPng(url string, dest string) error {
	w, h := 512, 512

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	icon, _ := oksvg.ReadIconStream(response.Body)
	icon.SetTarget(0, 0, float64(w), float64(h))

	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, rgba)
	if err != nil {
		return err
	}
	return err
}

func ConvertSvgToPngWithChrome(url string, dest string) error {
	dir := filepath.Dir(dest)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	chrome := svg2png.NewChrome().SetHeight(600).SetWith(600)
	if err := chrome.Screenshoot(url, dest); err != nil {
		return err
	}
	return nil
}
