package updater_test

import (
	"os"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/updater"
)

func Example() {
	multiStep := updater.NewMultistep(updater.Write(os.Stdout), updater.WithBarWidth(20))
	step1 := multiStep.AddStep(100)
	step2 := multiStep.AddStep(100)
	step3 := multiStep.AddStep(100)

	step1.Start("Running..")
	step1.Complete("Complete")
	step3.Start("Preparing!")
	step3.Complete("Complete!")
	step2.UpdateStatus(prettyprogress.Downloading, "Downloading..")
	step2.UpdateProgress(prettyprogress.Uploading, "Uploading..", 40)
	step2.Fail(":(")
}

func ExampleStep_Bar() {
	step := updater.NewStep(100, 5, updater.Writeln(os.Stdout))
	step.UpdateStatus(prettyprogress.Running, "Preparing")

	bar := step.Bar(prettyprogress.Downloading, "Downloading")
	bar.UpdateProgress(20)
	bar.UpdateProgress(80)
	bar.UpdateProgress(100)
	// Output:  ►  Preparing
	//  ↡  Downloading   [█    ]
	//  ↡  Downloading   [████ ]
	//  ↡  Downloading   [█████]
}
