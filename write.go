package prettyprogress

import "io"

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

		if f, ok := w.(interface{ Flush() error }); ok {
			f.Flush()
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

		if f, ok := w.(interface{ Flush() error }); ok {
			f.Flush()
		}
	}
}
