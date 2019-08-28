package main

import (
	"time"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/updater"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	multiStep := updater.NewMultistep(
		updater.Writeln(w),
		updater.WithBulletColor(
			prettyprogress.Complete, color.New(color.FgGreen).Sprint,
		),
	)

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

func doSomethingWithProgress(b interface{ UpdateProgress(int) }) {
	for i := 0; i <= 100; i++ {
		b.UpdateProgress(i)
		time.Sleep(5 * time.Millisecond)
	}
}
