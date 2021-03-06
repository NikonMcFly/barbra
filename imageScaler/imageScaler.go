package imageScaler

import (
	"errors"
	"image"
	"image/png"
	"math"
	"os"

	"github.com/nfnt/resize"
	"golang.org/x/exp/shiny/unit"
	"golang.org/x/image/math/fixed"
)

// NewScale ...
func NewScale(img image.Image, scale *Scale) (image.Image, error) {

	if scale.Line.Start.X == 0 && scale.Line.End.X == 0 && scale.Line.Start.Y == 0 && scale.Line.End.Y == 0 {
		return nil, errors.New("Please select a refrence line on the image")
	}

	if scale.getAxis() != multiAxis {

		mutiplyer := scale.mutiplyer()

		if scale.getAxis() == xAxis {
			xLength := float64(img.Bounds().Dx()) / mutiplyer
			return resize.Resize(uint(xLength), 0, img, resize.Lanczos3), nil
		} else if scale.getAxis() == yAxis {
			yLength := float64(img.Bounds().Dy()) / scale.mutiplyer()
			return resize.Resize(0, uint(yLength), img, resize.Lanczos3), nil
		}

	} else {
		// MultiAxis
		c := unit.Converter(Default)
		knownlength := c.Convert(scale.knownLength(), unit.Px)
		mutiplyer := math.Abs(scale.getHypotenusePixels().F / knownlength.F)
		xLength := float64(img.Bounds().Dx()) / mutiplyer

		return resize.Resize(uint(xLength), 0, img, resize.Lanczos3), nil
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

const defaultDPI = 72.0

// Default ...
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
func (t *Theme) getDPI() float64 {
	if t != nil && t.DPI != 0 {
		return t.DPI
	}
	return defaultDPI
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
		return t.getDPI()
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
