package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/updater"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	watcher := func(s string) {
		fmt.Fprintln(w, s)
		w.Flush()
	}

	multiStep := updater.NewMultistep(20, watcher)
	step1 := multiStep.AddStep(100)
	step2 := multiStep.AddStep(100)
	step3 := multiStep.AddStep(100)

	step1.Start("Running..")
	step2.Start("Running..")
	step3.Start("Running..")

	ch := make(chan struct{})
	go func() {
		time.Sleep(200 * time.Millisecond)
		doSomethingWithProgress(step1.Bar(prettyprogress.Downloading, "Downloading.."))
		step1.Complete("Done-zo")
		close(ch)
	}()

	doSomethingWithProgress(step2.Bar(prettyprogress.Uploading, "Uploading.."))
	step2.Complete("Done-zo")
	step3.Complete("Complete!")

	<-ch
}

func doSomethingWithProgress(b updater.ProgressUpdater) {
	for i := 0; i <= 100; i++ {
		b.UpdateProgress(i)
		time.Sleep(5 * time.Millisecond)
	}
}

func doSomethingWithProgressAndStatus(b updater.StatusUpdater) {
	for i := 0; i <= 100; i++ {
		b.UpdateProgress(prettyprogress.Downloading, fmt.Sprintf("Progressing %d", i), i)
		time.Sleep(5 * time.Millisecond)
	}
}
