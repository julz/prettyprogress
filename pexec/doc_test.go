package pexec_test

import (
	"os"
	"os/exec"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/pexec"
)

func Example() {
	steps := prettyprogress.NewMultistep(prettyprogress.Write(os.Stdout))

	step1 := pexec.Wrap(exec.Command("echo step1"), steps.AddStep("", 0))
	step2 := pexec.Wrap(exec.Command("echo step2"), steps.AddStep("", 0))
	step3 := pexec.Wrap(exec.Command("echo step3"), steps.AddStep("", 0))

	step1.Run()
	step2.Run()
	step3.Run()

	// Output:
	//
	//     Run 'echo step1'
	//
	//     Run 'echo step1'
	//     Run 'echo step2'
	//
	//     Run 'echo step1'
	//     Run 'echo step2'
	//     Run 'echo step3'
	//
	//  ►  Running 'echo step1'..
	//     Run 'echo step2'
	//     Run 'echo step3'
	//
	//  ✓  Finished 'echo step1'
	//     Run 'echo step2'
	//     Run 'echo step3'
	//
	//  ✓  Finished 'echo step1'
	//  ►  Running 'echo step2'..
	//     Run 'echo step3'
	//
	//  ✓  Finished 'echo step1'
	//  ✓  Finished 'echo step2'
	//     Run 'echo step3'
	//
	//  ✓  Finished 'echo step1'
	//  ✓  Finished 'echo step2'
	//  ►  Running 'echo step3'..
	//
	//  ✓  Finished 'echo step1'
	//  ✓  Finished 'echo step2'
	//  ✓  Finished 'echo step3'
}
