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

	for i := 0; i <= 150; i++ {
		buildBullet := prettyprogress.Running
		if i > 24 {
			buildBullet = prettyprogress.Complete
		}

		downloadBullet := prettyprogress.Downloading
		if i >= 100 {
			downloadBullet = prettyprogress.Complete
		}

		scanBullet := prettyprogress.Running
		if i >= 80 {
			scanBullet = prettyprogress.Complete
		}

		startBullet := prettyprogress.Future
		startingText := "Waiting to Start"
		if i > 100 {
			startBullet = prettyprogress.Running
			startingText = "Running.."
		}

		progress := 100
		if i < 100 {
			progress = i
		}

		if i >= 150 {
			startBullet = prettyprogress.Complete
			startingText = "Done"
		}

		fmt.Fprint(writer, prettyprogress.Steps{
			{
				Name:   "Building..",
				Bullet: buildBullet,
			},
			{
				Name:   "Downloading..",
				Bullet: downloadBullet,
				Bar:    prettyprogress.NewBar(progress, 100).String(),
			},
			{
				Name:   "Scanning..",
				Bullet: scanBullet,
			},
			{
				Name:   startingText,
				Bullet: startBullet,
			},
		})

		time.Sleep(time.Millisecond * 5)
		writer.Flush()
	}
}
