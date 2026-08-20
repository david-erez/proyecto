package converters

import (
	"image"
	"strings"
)

// El set de caracteres, de más "oscuro" (denso) a más "claro" (vacío).
const asciiChars = "@%#*+=-:. "

type AsciiConverter struct {
	Width int
}

func (a AsciiConverter) Convert(img image.Image) (string, error) {
	bounds := img.Bounds()
	origWidth := bounds.Dx()
	origHeidh := bounds.Dy()

	// calcular nueva altura manteniendo la relación de aspecto (aprox) y un factor vertical
	if origWidth == 0 || a.Width <= 0 {
		return "", nil
	}
	newHeidh := int(float64(origHeidh) * float64(a.Width) / float64(origWidth) * 0.5)

	var render strings.Builder

	for y := 0; y < newHeidh; y++ {
		for x := 0; x < a.Width; x++ {
			//se mapea el esultado a la coordenada original
			origX := x * origWidth / a.Width
			origY := y * origHeidh / newHeidh

			r, g, b, _ := img.At(bounds.Min.X+origX, bounds.Min.Y+origY).RGBA()

			// luminancia aproximada (0..255)
			gray := (0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))

			// mapear gray (0..255) al rango de asciiChars (0..len-1)
			idx := int((gray / 255.0) * float64(len(asciiChars)-1))
			if idx < 0 {
				idx = 0
			}
			if idx > len(asciiChars)-1 {
				idx = len(asciiChars) - 1
			}

			render.WriteByte(asciiChars[idx])
		}
		render.WriteByte('\n')
	}
	return render.String(), nil
}
