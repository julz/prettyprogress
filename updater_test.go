package prettyprogress_test

import (
	"testing"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/ui"
	"gotest.tools/assert"
)

func TestProgressUpdater(t *testing.T) {
	var recieved []string
	bar := prettyprogress.NewBar(100, 20, func(b string) {
		recieved = append(recieved, b)
	})

	ch := make(chan struct{})
	go func() {
		bar.Update(10)
		bar.Update(5)
		bar.Update(15)
		close(ch)
	}()

	<-ch

	assert.DeepEqual(t, recieved, []string{
		ui.NewBarWithWidth(10, 100, 20).String(),
		ui.NewBarWithWidth(5, 100, 20).String(),
		ui.NewBarWithWidth(15, 100, 20).String(),
	})
}

func TestStatusUpdater(t *testing.T) {
	var recieved []string
	step := prettyprogress.NewStep(10, 5, func(b string) {
		recieved = append(recieved, b)
	})

	ch := make(chan struct{})
	go func() {
		step.Update(ui.Running, "Hello")
		step.UpdateWithProgress(ui.Complete, "Done", 4)
		close(ch)
	}()

	<-ch

	assert.DeepEqual(t, recieved, []string{
		ui.Step{
			Bullet: ui.Running,
			Name:   "Hello",
			Bar:    "",
		}.String(),
		ui.Step{
			Bullet: ui.Complete,
			Name:   "Done",
			Bar:    ui.Bar{Progress: 4, Total: 10, Width: 5}.String(),
		}.String(),
	})
}
