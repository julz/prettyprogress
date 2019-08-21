package prettyprogress

import "math"

// Bar is a simple progress bar struct that knows how to String() itself in a pretty way
type Bar struct {
	Progress int
	Total    int
	Width    int
}

const defaultWidth int = 20

// NewBar creates a new Bar with a default width of 20 characters
func NewBar(progress, total int) Bar {
	return Bar{Progress: progress, Total: total, Width: defaultWidth}
}

// NewBarWithWidth creates a new Bar with the given progress, total and width
func NewBarWithWidth(progress, total, width int) Bar {
	return Bar{Progress: progress, Total: total, Width: width}
}

// String stringifies the Bar to a nice-looking unicode string
func (b Bar) String() string {
	s := "["
	for i := 0; i < b.Width; i++ {
		if float64(i) < math.Floor(float64(b.Progress)*(float64(b.Width)/float64(b.Total))) {
			s += "█"
		} else if float64(i) < float64(b.Progress)*(float64(b.Width)/float64(b.Total)) {
			s += "▌"
		} else {
			s += " "
		}
	}
	s += "]"

	return s
}
