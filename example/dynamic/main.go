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

	multiStep := dynamic.NewMultistepUpdater(20, watcher)
	step1 := multiStep.AddStep(100)
	step2 := multiStep.AddStep(100)
	step3 := multiStep.AddStep(100)

	step1.UpdateStatus(prettyprogress.Running, "Running..")
	step2.UpdateStatus(prettyprogress.Running, "Running..")
	step3.UpdateStatus(prettyprogress.Running, "Running..")

	doSomethingWithProgress(step2.Bar(prettyprogress.Downloading, "Downloading.."))
	step2.UpdateStatus(prettyprogress.Complete, "Done-zo")

	step3.UpdateStatus(prettyprogress.Complete, "Done-zo")

	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		step1.UpdateStatus(prettyprogress.Complete, "Done-zo")
		close(ch)
	}()

	<-ch
}

func doSomethingWithProgress(b dynamic.ProgressUpdater) {
	for i := 0; i <= 100; i++ {
		b.UpdateProgress(i)
		time.Sleep(5 * time.Millisecond)
	}
}

func doSomethingWithProgressAndStatus(b dynamic.StatusUpdater) {
	for i := 0; i <= 100; i++ {
		b.UpdateProgress(prettyprogress.Downloading, fmt.Sprintf("Progressing %d", i), i)
		time.Sleep(5 * time.Millisecond)
	}
}
