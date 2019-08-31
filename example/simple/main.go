package main

import (
	"fmt"
	"math"
	"time"

	"github.com/gookit/color"
	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress/ui"
)

func main() {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	for i := 0; i < 100; i++ {
		fmt.Fprint(writer, ui.Steps{
			{
				Name:            "Building..",
				Bullet:          ui.Complete,
				BulletColorFunc: color.New(color.FgGreen).Render,
			},
			{
				Name:   "Downloading..",
				Bullet: ui.Downloading,
				Bar:    ui.NewBar(i, 100).String(),
			},
			{
				Name:   "Scanning..",
				Bullet: ui.Running,
				Bar:    ui.NewBar(int(math.Min(float64(i), 50)), 50).String(),
			},
			{
				Name:   "Waiting to Start..",
				Bullet: ui.Future,
			},
		})

		time.Sleep(time.Millisecond * 5)
		writer.Flush()
	}
}
