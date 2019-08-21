package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
)

func main() {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	for i := 0; i < 100; i++ {
		fmt.Fprint(writer, prettyprogress.Plan{
			{
				Name:   "Building..",
				Bullet: prettyprogress.Complete,
			},
			{
				Name:   "Downloading..",
				Bullet: prettyprogress.Downloading,
				Bar:    prettyprogress.NewBar(i, 100),
			},
			{
				Name:   "Scanning..",
				Bullet: prettyprogress.Running,
			},
			{
				Name:   "Waiting to Start..",
				Bullet: prettyprogress.Future,
			},
		})

		time.Sleep(time.Millisecond * 5)
		writer.Flush()
	}
}
