package prettyprogress_test

import (
	"fmt"
	"os"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/ui"
)

func Example() {
	multiStep := prettyprogress.NewMultistep(
		func(s string) {
			fmt.Println("---")
			fmt.Println(s)
		},
		prettyprogress.WithBarWidth(5),
		prettyprogress.WithBarLabel(ui.PercentageLabel),
	)

	step1 := multiStep.AddStep(prettyprogress.WithStatus("Prepare.."))
	step2 := multiStep.AddStep(prettyprogress.WithStatus("Download XYZ.."))

	step1.Complete("Prepared")

	bar := step2.Bar(ui.Downloading, "Downloading..")
	bar.Update(80)

	step2.Fail("Download Failed")

	// Output: ---
	//
	//     Prepare..
	//
	// ---
	//
	//     Prepare..
	//     Download XYZ..
	//
	// ---
	//
	//  ✓  Prepared
	//     Download XYZ..
	//
	// ---
	//
	//  ✓  Prepared
	//  ↡  Downloading..   [████ ] 80%
	//
	// ---
	//
	//  ✓  Prepared
	//  ✗  Download Failed
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
