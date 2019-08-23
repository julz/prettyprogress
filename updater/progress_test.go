package updater_test

import (
	"testing"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/updater"
	"gotest.tools/assert"
)

func TestMultistep(t *testing.T) {
	var recieved []string
	steps := updater.NewMultistep(20, func(s string) {
		recieved = append(recieved, s)
	})

	step1 := steps.AddStep(100)
	step2 := steps.AddStep(10)

	step1.UpdateStatus(prettyprogress.Running, "hello")
	step2.UpdateStatus(prettyprogress.Complete, "bye")
	step1.UpdateProgress(prettyprogress.Downloading, "updated", 12)

	assert.DeepEqual(t, recieved, []string{
		prettyprogress.Steps{
			{
				Bullet: prettyprogress.Running,
				Name:   "hello",
			},
			{
				Bullet: prettyprogress.Future,
				Name:   "",
			},
		}.String(),
		prettyprogress.Steps{
			{
				Bullet: prettyprogress.Running,
				Name:   "hello",
			},
			{
				Bullet: prettyprogress.Complete,
				Name:   "bye",
			},
		}.String(),
		prettyprogress.Steps{
			{
				Bullet: prettyprogress.Downloading,
				Name:   "updated",
				Bar:    prettyprogress.NewBarWithWidth(12, 100, 20).String(),
			},
			{
				Bullet: prettyprogress.Complete,
				Name:   "bye",
			},
		}.String(),
	})
}