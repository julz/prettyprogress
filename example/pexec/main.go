package main

import (
	"os/exec"

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

	steps := prettyprogress.NewMultistep(
		prettyprogress.Writeln(w),
		prettyprogress.WithBulletColor(
			ui.Complete, color.New(color.FgGreen).Sprint,
		),
	)

	step1 := pexec.Wrap(exec.Command("bash", "-c", "sleep 2; echo hello"), steps.AddStep("", 0))
	step2 := pexec.Wrap(exec.Command("bash", "-c", "sleep 3; echo hello"), steps.AddStep("", 0))
	step3 := pexec.Wrap(exec.Command("bash", "-c", "sleep 1; echo hello"), steps.AddStep("", 0))

	go step3.Run()
	step1.Run()
	step2.Run()
}
