package ui_test

import (
	"fmt"

	"github.com/julz/prettyprogress/ui"
)

func Example() {
	fmt.Printf("\n%s", ui.Steps{
		{
			Name:   "Building..",
			Bullet: ui.Complete,
		},
		{
			Name:   "Downloading..",
			Bullet: ui.Downloading,
			Bar:    ui.NewBar(10, 100).String(),
		},
		{
			Name:   "Scanning..",
			Bullet: ui.Running,
			Bar:    ui.NewBar(20, 100).String(),
		},
		{
			Name:   "Waiting to Start..",
			Bullet: ui.Future,
		},
	})

	// Output:
	//  ✓  Building..
	//  ↡  Downloading..        [██                  ]
	//  ►  Scanning..           [████                ]
	//     Waiting to Start..
}

func ExampleStep() {
	fmt.Printf("%s", ui.Step{
		Bullet: ui.Downloading,
		Name:   "Downloading..",
		Bar:    "[██   ]",
	})

	// Output: ↡  Downloading..   [██   ]
}
func ExampleBar() {
	fmt.Printf("%s", ui.Bar{
		Width:    5,
		Total:    100,
		Progress: 40,
	})

	// Output: [██   ]
}
