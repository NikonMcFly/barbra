package imageScaler

import (
	"math"

	"golang.org/x/exp/shiny/unit"
)

const (
	// XAxis ...
	XAxis = "x"
	// YAxis ...
	YAxis = "y"
)

type point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type line struct {
	Start *point `json:"start"`
	End   *point `json:"end"`
}

// Scale ...
type Scale struct {
	Line   *line   `json:"line"`
	Length float64 `json:"length"`
	Axis   string
}

func newPoint() *point {
	return &point{
		X: 0,
		Y: 0,
	}
}

func newLine() *line {
	return &line{
		Start: newPoint(),
		End:   newPoint(),
	}
}

// NewTransformation initializes a Scale struct with default settings
func NewTransformation() *Scale {
	return &Scale{
		Line:   newLine(),
		Length: 0,
	}
}

// Mutiplyer ...
func (s *Scale) Mutiplyer() float64 {

	// if !s.isSingleAxis() {
	// 	return s.getHypotenusePixels().F / s.KnownLength().F
	// }

	c := unit.Converter(Default)
	knownPixelLength := c.Convert(s.KnownLength(), unit.Px)
	return s.Pixels().F / knownPixelLength.F
}

// Pixels  returns the end point of the line
func (s *Scale) Pixels() unit.Value {

	if s.Axis == XAxis || !s.isSingleAxis() {
		return unit.Pixels(s.Line.End.X - s.Line.Start.X)
	}

	return unit.Pixels(s.Line.End.Y - s.Line.Start.Y)

}

// KnownLength returns known length of the Li
func (s *Scale) KnownLength() unit.Value {

	return unit.Inches(s.Length)
}

func (s *Scale) isSingleAxis() bool {
	if s.Line.Start.X == s.Line.End.X {
		return true
	}

	if s.Line.Start.Y == s.Line.End.Y {
		return true
	}

	return false
}

func (s *Scale) getHypotenusePixels() unit.Value {

	x2Minusx1 := (s.Line.End.X - s.Line.Start.X)

	x2Minusx1Squared := x2Minusx1 * x2Minusx1

	y2Minusy1 := (s.Line.End.Y - s.Line.Start.Y)

	y2Minusy1Squared := y2Minusy1 * y2Minusy1

	hypotenuse := math.Sqrt(x2Minusx1Squared + y2Minusy1Squared)

	return unit.Pixels(hypotenuse)
}
