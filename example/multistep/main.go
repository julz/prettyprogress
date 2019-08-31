package main

import (
	"time"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/ui"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	multiStep := prettyprogress.NewMultistep(
		prettyprogress.Writeln(w),
		prettyprogress.WithBulletColor(
			ui.Complete, color.New(color.FgGreen).Sprint,
		),
	)

	step1 := multiStep.AddStep(prettyprogress.WithBarTotal(1000))
	step2 := multiStep.AddStep(prettyprogress.WithBarTotal(1000))
	step3 := multiStep.AddStep()

	step1.Start("Running..")
	step2.Start("Running..")
	step3.Start("Running..")

	ch := make(chan struct{})
	go func() {
		time.Sleep(500 * time.Millisecond)
		doSomethingWithProgress(step1.Bar(ui.Downloading, "Downloading.."))
		step1.Complete("Done-zo")
		close(ch)
	}()

	doSomethingWithProgress(step2.Bar(ui.Uploading, "Uploading.."))
	step2.Complete("Done-zo")
	step3.Complete("Complete!")

	<-ch
}

func doSomethingWithProgress(bar interface{ Update(int) }) {
	for i := 0; i <= 1010; i++ {
		bar.Update(i)
		time.Sleep(2 * time.Millisecond)
	}
}
