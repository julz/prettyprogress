package main

import (
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/pexec"
	"github.com/julz/prettyprogress/ui"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	bullets := ui.AnimatedBulletSet
	bullets.Running = bullets.Running.WithColor(color.New(color.FgGreen))

	steps := prettyprogress.NewMultistep(
		prettyprogress.Writeln(w),
		prettyprogress.WithAnimationFrameTicker(time.NewTicker(200*time.Millisecond).C),
		prettyprogress.WithBullets(
			bullets,
		),
	)

	step1 := pexec.Wrap(exec.Command("bash", "-c", "sleep 2; echo hello"), steps.AddStep("", 0))
	step2 := pexec.Wrap(exec.Command("bash", "-c", "sleep 3; echo hello"), steps.AddStep("", 0))
	step3 := pexec.Wrap(exec.Command("bash", "-c", "sleep 1; echo hello"), steps.AddStep("", 0))

	go step3.Run()
	step1.Run()
	step2.Run()
}
