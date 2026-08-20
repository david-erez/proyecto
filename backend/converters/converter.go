package converters

import (
	"fmt"
	"image"
)

type ImageConverter interface {
	Convert(img image.Image) (string, error)
}

func NewConverter(typo string) (ImageConverter, error) {
	switch typo {
	case "ascii":
		return AsciiConverter{Width: 100}, nil
	case "pixelart":
		return PixelArtConverter{BlockSize: 8}, nil
	default:
		return nil, fmt.Errorf("tipo de conversión no soportado: %s", typo)
	}
}
