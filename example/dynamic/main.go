package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/dynamic"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	watcher := func(s string) {
		fmt.Fprintln(w, s)
		w.Flush()
	}

	bar := dynamic.NewProgressUpdater(100, 20, watcher)
	doSomethingWithProgress(bar)

	step := dynamic.NewStatusUpdater(100, 20, watcher)
	doSomethingWithProgressAndStatus(step)
}

func doSomethingWithProgress(b dynamic.ProgressUpdater) {
	for i := 0; i <= 100; i++ {
		b.Update(i)
		time.Sleep(5 * time.Millisecond)
	}
}

func doSomethingWithProgressAndStatus(b dynamic.StatusUpdater) {
	for i := 0; i <= 100; i++ {
		b.UpdateProgress(prettyprogress.Downloading, fmt.Sprintf("Progressing %d", i), i)
		time.Sleep(5 * time.Millisecond)
	}
}
