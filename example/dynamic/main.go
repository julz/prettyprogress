package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress/dynamic"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	bar := dynamic.NewProgressUpdater(100, 20, func(b string) {
		fmt.Fprintln(w, b)
		w.Flush()
	})

	doSomethingWithProgress(bar)
}

func doSomethingWithProgress(b dynamic.ProgressUpdater) {
	for i := 0; i <= 100; i++ {
		b.Update(i)
		time.Sleep(5 * time.Millisecond)
	}
}
