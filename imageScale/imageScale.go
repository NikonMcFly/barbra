package imageScale

import (
	"errors"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
	"golang.org/x/exp/shiny/unit"
	"golang.org/x/image/math/fixed"
)

func main() {
	img, err := GetPng("./images/University of Houston Logo.png")
	if err != nil {
		log.Fatal(err)
	}

	pixelScale := unit.Pixels(144)
	knownLength := unit.Inches(4)

	m, _ := ScaleImage(img, pixelScale, knownLength, "x")

	out, err := os.Create("test_reszied.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	png.Encode(out, m)

}

// ScaleImage ...
func ScaleImage(img image.Image, pixelScale unit.Value, knownLength unit.Value, axis string) (image.Image, error) {

	c := unit.Converter(Default)
	knownPixelLength := c.Convert(knownLength, unit.Px)

	if pixelScale.F == 0 {
		if axis == "x" {
			return resize.Resize(uint(knownPixelLength.F), 0, img, resize.Lanczos3), nil
		} else if axis == "y" {
			return resize.Resize(0, uint(knownPixelLength.F), img, resize.Lanczos3), nil
		}
	} else {

		// The pixelScale is either in the x or the y axis because it is a vertical or horizontal line
		// I beleve that we cold handle lines drawn at any angle using the Pythagorean theorem

		mutiplyer := (pixelScale.F / knownPixelLength.F)

		if axis == "x" {
			xLength := float64(img.Bounds().Dx()) / mutiplyer
			return resize.Resize(uint(xLength), 0, img, resize.Lanczos3), nil
		} else if axis == "y" {
			yLength := float64(img.Bounds().Dy()) / mutiplyer
			return resize.Resize(0, uint(yLength), img, resize.Lanczos3), nil
		}
	}

	return nil, errors.New("axis is not supported")
}

// GetPng ...
func GetPng(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, err
}

// does not provide a DPI value.
const DefaultDPI = 72.0

var Default *Theme

// Theme ..
type Theme struct {
	DPI float64
}

// Pixels implements the unit.Converter interface.
func (t *Theme) Pixels(v unit.Value) fixed.Int26_6 {
	c := t.Convert(v, unit.Px)
	return fixed.Int26_6(c.F * 64)
}

// Convert implements the unit.Converter interface.
func (t *Theme) Convert(v unit.Value, to unit.Unit) unit.Value {
	if v.U == to {
		return v
	}
	return unit.Value{
		F: v.F * t.pixelsPer(v.U) / t.pixelsPer(to),
		U: to,
	}
}

// GetDPI returns the theme's DPI, or the default DPI if the field value is
// zero.
func (t *Theme) GetDPI() float64 {
	if t != nil && t.DPI != 0 {
		return t.DPI
	}
	return DefaultDPI
}

// pixelsPer returns the number of pixels in the unit u.
func (t *Theme) pixelsPer(u unit.Unit) float64 {
	switch u {
	// case unit.Px:
	// 	return 1
	// case unit.Dp:
	// 	return t.GetDPI() / unit.DensityIndependentPixelsPerInch
	// case unit.Pt:
	// 	return t.GetDPI() / unit.PointsPerInch
	// case unit.Mm:
	// 	return t.GetDPI() / unit.MillimetresPerInch
	case unit.In:
		return t.GetDPI()
	}
	// f := t.AcquireFontFace(FontFaceOptions{})
	// defer t.ReleaseFontFace(FontFaceOptions{}, f)
	// The 64 is because Height is in 26.6 fixed-point units.
	// h := float64(f.Metrics().Height) / 64
	// switch u {
	// case unit.Em:
	// 	return h
	// case unit.Ex:
	// 	return h / 2
	// case unit.Ch:
	// 	if advance, ok := f.GlyphAdvance('0'); ok {
	// 		return float64(advance) / 64
	// 	}
	// 	return h / 2
	// }
	return 1
}
