package updater_test

import (
	"os"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/updater"
)

func Example() {
	multiStep := updater.NewMultistep(updater.Write(os.Stdout), updater.WithBarWidth(20))
	step1 := multiStep.AddStep()
	step2 := multiStep.AddStep()
	step3 := multiStep.AddStep()

	step1.Start("Running..")
	step1.Complete("Complete")
	step3.Start("Preparing!")
	step3.Complete("Complete!")
	step2.Update(prettyprogress.Downloading, "Downloading..")
	step2.UpdateWithProgress(prettyprogress.Uploading, "Uploading..", 40)
	step2.Fail(":(")
}

func ExampleBar() {
	bar := updater.NewBar(100, 5, updater.Writeln(os.Stdout))
	bar.Update(20)
	bar.Update(40)
	bar.Update(80)

	// Output:  [█    ]
	// [██   ]
	// [████ ]
}

func ExampleStep() {
	step := updater.NewStep(100, 5, updater.Writeln(os.Stdout))
	step.Start("Starting..")
	step.UpdateWithProgress(prettyprogress.Downloading, "Downloading", 20)
	step.UpdateWithProgress(prettyprogress.Downloading, "Downloading", 80)
	step.Complete("Done!")

	// Output:  ►  Starting..
	//  ↡  Downloading   [█    ]
	//  ↡  Downloading   [████ ]
	//  ✓  Done!
}

func ExampleStep_Bar() {
	step := updater.NewStep(100, 5, updater.Writeln(os.Stdout))
	step.Update(prettyprogress.Running, "Preparing")

	bar := step.Bar(prettyprogress.Downloading, "Downloading")
	bar.Update(20)
	bar.Update(80)
	bar.Update(100)
	// Output:  ►  Preparing
	//  ↡  Downloading   [█    ]
	//  ↡  Downloading   [████ ]
	//  ↡  Downloading   [█████]
}
