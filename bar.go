package prettyprogress

import "github.com/julz/prettyprogress/ui"

// Bar is an updater that calls its watcher with the String() version of a
// ui.Bar whenever its UpdateProgress method is called.
type Bar struct {
	total int
	width int

	label ui.LabelFunc

	printFunc PrintFunc
}

// NewBar returns a new Bar object that can be updated, when the UpdateProgress
// method is called, the passed Watcher is called with the new state. This can
// be used to print the bar out, either directly to standard out or with some
// terminal magic to do animation.
func NewBar(total, width int, w PrintFunc) *Bar {
	return &Bar{
		total:     total,
		width:     width,
		printFunc: w,
	}
}

// Update calls the Bar's watcher (configured in NewBar) with the new
// state. Generally this will cause the new bar to be printed out to the user.
func (b *Bar) Update(progress int) {
	b.printFunc(ui.Bar{Progress: progress, Width: b.width, Total: b.total, LabelFunc: b.label}.String())
}
