package imageScaler

import "golang.org/x/exp/shiny/unit"

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
	c := unit.Converter(Default)
	knownPixelLength := c.Convert(s.KnownLength(), unit.Px)
	return s.Pixels().F / knownPixelLength.F
}

// Pixels  returns the end point of the line
func (s *Scale) Pixels() unit.Value {
	if s.Axis == "x" {
		return unit.Pixels(s.Line.End.X - s.Line.Start.X)
	}
	return unit.Pixels(s.Line.End.Y - s.Line.Start.Y)
}

// KnownLength returns known length of the Li
func (s *Scale) KnownLength() unit.Value {
	return unit.Inches(s.Length)
}
