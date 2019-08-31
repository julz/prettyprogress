package ui

import (
	"math"
	"strings"
)

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
	fractions := []string{
		" ",
		"▏",
		"▎",
		"▍",
		"▌",
		"▋",
		"▊",
		"▉",
		"█",
	}

	progress := b.Progress
	if b.Progress > b.Total {
		progress = b.Total
	}

	s := "["

	// scaledProgress is progress between 0 and width (rather than 0 and total)
	scaledProgress := float64(progress) * (float64(b.Width) / float64(b.Total))

	// split in to whole-sized cells and fractional part
	// we can paint all the whole-sized cells with █ and then
	// use a fractional unicode character for the fractional cell
	wholeCells, remainder := math.Modf(scaledProgress)

	s += strings.Repeat("█", int(wholeCells))

	// fill in the remainder if the bar isn't full yet
	if int(wholeCells) < b.Width {
		// convert fractional (0-1) remainder to 1-8 unicode characters
		// so we have greater resolution that the number of actual console characters
		s += fractions[int(math.Floor(remainder*8))]

		// fill the rest with spaces
		s += strings.Repeat(" ", b.Width-(int(wholeCells)+1))
	}

	s += "]"
	return s
}
