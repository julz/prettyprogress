package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/pexec"
	"github.com/julz/prettyprogress/ui"
)

//  e.g. `go run example/pexec/main.go "sleep 1" "echo hello" "sleep 2"`
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

	args := []string{}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	cmds := []*pexec.Cmd{}
	for _, a := range args {
		cmds = append(cmds, pexec.Wrapn(a, exec.Command("sh", "-c", a), steps.AddStep("", 0)))
	}

	for _, cmd := range cmds {
		cmd.Run()
	}
}
