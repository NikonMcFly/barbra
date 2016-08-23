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
	// MultiAxis
	MultiAxis = "xy"
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

	c := unit.Converter(Default)
	knownPixelLength := c.Convert(s.KnownLength(), unit.Px)

	if s.Pixels().F == 0 {
		return 1
	}

	return s.Pixels().F / knownPixelLength.F
}

// Pixels  returns the end point of the line
func (s *Scale) Pixels() unit.Value {

	if s.getAxis() == XAxis || s.getAxis() == MultiAxis {
		return unit.Pixels(s.Line.End.X - s.Line.Start.X)
	}

	return unit.Pixels(s.Line.End.Y - s.Line.Start.Y)

}

// KnownLength returns known length of the Li
func (s *Scale) KnownLength() unit.Value {

	return unit.Inches(s.Length)
}

func (s *Scale) isSingleAxis() bool {
	if s.getAxis() == MultiAxis {
		return false
	}
	return true
}

func (s *Scale) getAxis() string {

	if s.Line.Start.Y == s.Line.End.Y {
		return XAxis
	}

	if s.Line.Start.X == s.Line.End.X {
		return YAxis
	}

	return MultiAxis
}

func (s *Scale) getHypotenusePixels() unit.Value {

	x2Minusx1 := (s.Line.End.X - s.Line.Start.X)

	x2Minusx1Squared := x2Minusx1 * x2Minusx1

	y2Minusy1 := (s.Line.End.Y - s.Line.Start.Y)

	y2Minusy1Squared := y2Minusy1 * y2Minusy1

	hypotenuse := math.Sqrt(x2Minusx1Squared + y2Minusy1Squared)

	return unit.Pixels(hypotenuse)
}
