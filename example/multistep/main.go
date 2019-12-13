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

	multiStep := prettyprogress.NewFancyMultistep(
		w,
		prettyprogress.WithBarLabel(ui.PercentageLabel),
	)

	step1 := multiStep.AddStep("Download", 1000)
	step2 := multiStep.AddStep("Run Something", 0)
	step3 := multiStep.AddStep("Upload", 800)

	ch := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		doSomethingWithProgress(step1.Bar(ui.Downloading, "Downloading.."))
		step1.Complete("Download Done")
		close(ch)
	}()

	time.Sleep(1600 * time.Millisecond)
	step2.Start("Running Something..")

	failed := multiStep.AddStep("Waiting to fail", 0)
	doSomethingWithProgress(step3.Bar(ui.Uploading, "Uploading.."))
	step3.Complete("Upload Complete")

	time.Sleep(200 * time.Millisecond)
	step2.Complete("Some Action Done")

	time.Sleep(600 * time.Millisecond)
	failed.Start("Failing..")
	time.Sleep(400 * time.Millisecond)
	failed.Fail("Something failed!")

	cancel := multiStep.AddStep("", 0)
	cancel.Start("Cancelled")
	time.Sleep(500 * time.Millisecond)
	cancel.Update(ui.Complete, color.New(color.FgYellow).Sprint("Cancelled")) // override default colors

	<-ch
}

func doSomethingWithProgress(bar interface{ Update(int) }) {
	time.Sleep(300 * time.Millisecond)
	for i := 0; i <= 1000; i++ {
		bar.Update(i)
		time.Sleep(1 * time.Millisecond)
	}
}
