package converters

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
)

type PixelArtConverter struct {
	BlockSize int
}

func (p PixelArtConverter) Convert(img image.Image) (string, error) {

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	output := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += p.BlockSize {
		for x := bounds.Min.X; x < bounds.Max.X; x += p.BlockSize {
			avgColor := averageColor(img, x, y, p.BlockSize, width, height)

			for by := 0; by < p.BlockSize && y+by < bounds.Max.Y; by++ {
				for bx := 0; bx < p.BlockSize && x+bx < bounds.Max.X; bx++ {
					output.Set(x+bx, y+by, avgColor)
				}
			}
		}
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, output); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func averageColor(img image.Image, startX, startY, blockSize, width, height int) color.Color {
	var rSum, gSum, bSum, count uint32

	for y := startY; y < startY+blockSize && y < height; y++ {
		for x := startX; x < startX+blockSize && x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			rSum += r >> 8
			gSum += g >> 8
			bSum += b >> 8
			count++
		}
	}

	if count == 0 {
		return color.RGBA{0, 0, 0, 255}
	}

	return color.RGBA{
		R: uint8(rSum / count),
		G: uint8(gSum / count),
		B: uint8(bSum / count),
		A: 255,
	}
}
