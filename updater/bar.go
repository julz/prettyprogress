package updater

import (
	"io"

	"github.com/julz/prettyprogress"
)

// Watcher is a function that is called when new versions of a Bar, Step or Multistep
// are created as a result of calling UpdateX methods.
type Watcher func(s string)

// Bar is an updater that calls its watcher with the String() version of a
// prettyprogress.Bar whenever its UpdateProgress method is called.
type Bar struct {
	total int
	width int

	watcher Watcher
}

// NewBar returns a new Bar object that can be updated, when the UpdateProgress
// method is called, the passed Watcher is called with the new state. This can
// be used to print the bar out, either directly to standard out or with some
// terminal magic to do animation.
func NewBar(total, width int, w Watcher) *Bar {
	return &Bar{
		total:   total,
		width:   width,
		watcher: w,
	}
}

// UpdateProgress calls the Bar's watcher (configured in NewBar) with the new
// state. Generally this will cause the new bar to be printed out to the user.
func (b *Bar) UpdateProgress(progress int) {
	b.watcher(prettyprogress.Bar{Progress: progress, Width: b.width, Total: b.total}.String())
}

// Write creates a Watcher function that writes to a given io.Writer, useful for
// passing to one of the constructors in this package, e.g.
//
//   NewBar(100,100,Write(os.Stdout))
//
// Write will panic if it is unable to write to the underlying writer
func Write(w io.Writer) func(s string) {
	return func(s string) {
		_, err := w.Write([]byte(s))
		if err != nil {
			panic(err)
		}
	}
}

// Writeln creates a Watcher function that writes to a given io.Writer after
// appending a newline, useful for passing to one of the constructors in this
// package, e.g.
//
//   NewBar(100,100,Writeln(os.Stdout))
//
// Writeln will panic if it is unable to write to the underlying writer
func Writeln(w io.Writer) func(s string) {
	return func(s string) {
		_, err := w.Write([]byte(s + "\n"))
		if err != nil {
			panic(err)
		}
	}
}
