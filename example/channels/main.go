package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
)

func main() {
	downloadProgress := make(chan int)

	go func() {
		defer close(downloadProgress)

		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 5)
			downloadProgress <- i
		}
	}()

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	progress := 0
	downloading := true
	downloadStatus := prettyprogress.Downloading

	for {
		select {
		case progress, downloading = <-downloadProgress:
		}

		if !downloading {
			downloadStatus = prettyprogress.Complete
		}

		fmt.Fprint(writer, prettyprogress.Plan{
			{
				Name:   "Building..",
				Bullet: prettyprogress.Complete,
			},
			{
				Name:   "Downloading..",
				Bullet: downloadStatus,
				Bar: prettyprogress.Bar{
					Progress: progress,
					Total:    100,
					Width:    20,
				},
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

		if !downloading {
			break
		}
	}
}
