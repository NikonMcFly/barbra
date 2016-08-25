package imageScaler

import (
	"math"

	"golang.org/x/exp/shiny/unit"
)

const (
	xAxis     = "x"
	yAxis     = "y"
	multiAxis = "xy"
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

func (s *Scale) mutiplyer() float64 {

	c := unit.Converter(Default)
	knownPixelLength := c.Convert(s.knownLength(), unit.Px)

	if s.pixels().F == 0 {
		return 1
	}

	return s.pixels().F / knownPixelLength.F
}

func (s *Scale) pixels() unit.Value {

	if s.getAxis() == xAxis || s.getAxis() == multiAxis {
		return unit.Pixels(s.Line.End.X - s.Line.Start.X)
	}

	return unit.Pixels(s.Line.End.Y - s.Line.Start.Y)

}

func (s *Scale) knownLength() unit.Value {

	return unit.Inches(s.Length)
}

func (s *Scale) isSingleAxis() bool {
	if s.getAxis() == multiAxis {
		return false
	}
	return true
}

func (s *Scale) getAxis() string {

	if s.Line.Start.Y == s.Line.End.Y {
		return xAxis
	}

	if s.Line.Start.X == s.Line.End.X {
		return yAxis
	}

	return multiAxis
}

func (s *Scale) getHypotenusePixels() unit.Value {

	x2Minusx1 := (s.Line.End.X - s.Line.Start.X)

	x2Minusx1Squared := x2Minusx1 * x2Minusx1

	y2Minusy1 := (s.Line.End.Y - s.Line.Start.Y)

	y2Minusy1Squared := y2Minusy1 * y2Minusy1

	hypotenuse := math.Sqrt(x2Minusx1Squared + y2Minusy1Squared)

	absHypot := math.Abs(hypotenuse)

	return unit.Pixels(absHypot)
}
