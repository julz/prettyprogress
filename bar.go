package prettyprogress

import "math"

type Bar struct {
	Progress int
	Total    int
	Width    int
}

const defaultWidth int = 20

func NewBar(progress, total int) Bar {
	return Bar{Progress: progress, Total: total, Width: defaultWidth}
}

func NewBarWithWidth(progress, total, width int) Bar {
	return Bar{Progress: progress, Total: total, Width: width}
}

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
