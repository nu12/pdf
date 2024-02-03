/*
Copyright © 2024 nu12
*/

package model

import (
	"image"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/draw"
)

type PNG struct {
	Filename string
}

func (img PNG) Compress(p CompressionProfile) (io.Reader, error) {
	// Resize: https://stackoverflow.com/questions/22940724/go-resizing-images
	// Available methods are NearestNeighbor, ApproxBiLinear, BiLinear, CatmullRom
	pr, pw := io.Pipe()

	file, err := os.Open(img.Filename)
	if err != nil {
		return pr, err
	}
	src, err := png.Decode(file)
	if err != nil {
		return pr, err
	}

	dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/p.Value, src.Bounds().Max.Y/p.Value))

	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	go func() {
		_ = png.Encode(pw, dst)
		pw.Close()

	}()
	return pr, nil
}

func (PNG) GetType() string {
	return "png"
}

func (img PNG) GetFilename() string {
	return img.Filename
}
