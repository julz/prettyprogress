package prettyprogress_test

import (
	"fmt"

	"github.com/julz/prettyprogress"
)

func Example() {
	fmt.Printf("\n%s", prettyprogress.Steps{
		{
			Name:   "Building..",
			Bullet: prettyprogress.Complete,
		},
		{
			Name:   "Downloading..",
			Bullet: prettyprogress.Downloading,
			Bar:    prettyprogress.NewBar(10, 100).String(),
		},
		{
			Name:   "Scanning..",
			Bullet: prettyprogress.Running,
			Bar:    prettyprogress.NewBar(20, 100).String(),
		},
		{
			Name:   "Waiting to Start..",
			Bullet: prettyprogress.Future,
		},
	})
	// Output:
	//  ✓  Building..
	//  ↡  Downloading..        [██                  ]
	//  ►  Scanning..           [████                ]
	//     Waiting to Start..
}
