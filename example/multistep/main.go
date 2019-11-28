package main

import (
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/ui"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	multiStep := prettyprogress.NewFancyMultistep(
		prettyprogress.Writeln(w),
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

	doSomethingWithProgress(step3.Bar(ui.Uploading, "Uploading.."))
	step3.Complete("Upload Complete")

	time.Sleep(200 * time.Millisecond)
	step2.Complete("Some Action Done")

	<-ch
}

func doSomethingWithProgress(bar interface{ Update(int) }) {
	time.Sleep(300 * time.Millisecond)
	for i := 0; i <= 1000; i++ {
		bar.Update(i)
		time.Sleep(1 * time.Millisecond)
	}
}
