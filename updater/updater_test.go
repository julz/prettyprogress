package updater_test

import (
	"testing"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/updater"
	"gotest.tools/assert"
)

func TestProgressUpdater(t *testing.T) {
	var recieved []string
	updater := updater.NewBar(100, 20, func(b string) {
		recieved = append(recieved, b)
	})

	ch := make(chan struct{})
	go func() {
		updater.UpdateProgress(10)
		updater.UpdateProgress(5)
		updater.UpdateProgress(15)
		close(ch)
	}()

	<-ch

	assert.DeepEqual(t, recieved, []string{
		prettyprogress.NewBarWithWidth(10, 100, 20).String(),
		prettyprogress.NewBarWithWidth(5, 100, 20).String(),
		prettyprogress.NewBarWithWidth(15, 100, 20).String(),
	})
}

func TestStatusUpdater(t *testing.T) {
	var recieved []string
	updater := updater.NewStep(10, 5, func(b string) {
		recieved = append(recieved, b)
	})

	ch := make(chan struct{})
	go func() {
		updater.UpdateStatus(prettyprogress.Running, "Hello")
		updater.UpdateProgress(prettyprogress.Complete, "Done", 4)
		close(ch)
	}()

	<-ch

	assert.DeepEqual(t, recieved, []string{
		prettyprogress.Step{
			Bullet: prettyprogress.Running,
			Name:   "Hello",
			Bar:    "",
		}.String(),
		prettyprogress.Step{
			Bullet: prettyprogress.Complete,
			Name:   "Done",
			Bar:    prettyprogress.Bar{Progress: 4, Total: 10, Width: 5}.String(),
		}.String(),
	})
}
