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
		prettyprogress.WithBarLabel(ui.PercentageLabel),
	)

	step1 := multiStep.AddStep("Download", 1000)
	step2 := multiStep.AddStep("Upload", 800)
	step3 := multiStep.AddStep("Some Action", 0)

	step3.Start("Running Something..")

	ch := make(chan struct{})
	go func() {
		time.Sleep(400 * time.Millisecond)
		doSomethingWithProgress(step1.Bar(ui.Downloading, "Downloading.."))
		step1.Complete("Download done")
		close(ch)
	}()

	doSomethingWithProgress(step2.Bar(ui.Uploading, "Uploading.."))
	step2.Complete("Upload complete")

	time.Sleep(600 * time.Millisecond)
	step3.Complete("Some Action Done")

	<-ch
}

func doSomethingWithProgress(bar interface{ Update(int) }) {
	for i := 0; i <= 1010; i++ {
		bar.Update(i)
		time.Sleep(2 * time.Millisecond)
	}
}
