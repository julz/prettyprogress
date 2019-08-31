package prettyprogress_test

import (
	"os"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/ui"
)

func Example() {
	multiStep := prettyprogress.NewMultistep(prettyprogress.Write(os.Stdout), prettyprogress.WithBarWidth(20))
	step1 := multiStep.AddStep()
	step2 := multiStep.AddStep()
	step3 := multiStep.AddStep()

	step1.Start("Running..")
	step1.Complete("Complete")
	step3.Start("Preparing!")
	step3.Complete("Complete!")
	step2.Update(ui.Downloading, "Downloading..")
	step2.UpdateWithProgress(ui.Uploading, "Uploading..", 40)
	step2.Fail(":(")
}

func ExampleBar() {
	bar := prettyprogress.NewBar(100, 5, prettyprogress.Writeln(os.Stdout))
	bar.Update(20)
	bar.Update(40)
	bar.Update(80)

	// Output:  [█    ]
	// [██   ]
	// [████ ]
}

func ExampleStep() {
	step := prettyprogress.NewStep(100, 5, prettyprogress.Writeln(os.Stdout))
	step.Start("Starting..")
	step.UpdateWithProgress(ui.Downloading, "Downloading", 20)
	step.UpdateWithProgress(ui.Downloading, "Downloading", 80)
	step.Complete("Done!")

	// Output:  ►  Starting..
	//  ↡  Downloading   [█    ]
	//  ↡  Downloading   [████ ]
	//  ✓  Done!
}

func ExampleStep_Bar() {
	step := prettyprogress.NewStep(100, 5, prettyprogress.Writeln(os.Stdout))
	step.Update(ui.Running, "Preparing")

	bar := step.Bar(ui.Downloading, "Downloading")
	bar.Update(20)
	bar.Update(80)
	bar.Update(100)
	// Output:  ►  Preparing
	//  ↡  Downloading   [█    ]
	//  ↡  Downloading   [████ ]
	//  ↡  Downloading   [█████]
}
