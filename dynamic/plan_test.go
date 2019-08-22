package dynamic_test

import (
	"testing"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/dynamic"
	"gotest.tools/assert"
)

func TestMultistep(t *testing.T) {
	var recieved []string
	steps := dynamic.NewMultistepUpdater(20, func(s string) {
		recieved = append(recieved, s)
	})

	step1 := steps.AddStep(100)
	step2 := steps.AddStep(10)

	step1.Update(prettyprogress.Running, "hello")
	step2.Update(prettyprogress.Complete, "bye")
	step1.UpdateProgress(prettyprogress.Downloading, "updated", 12)

	assert.DeepEqual(t, recieved, []string{
		prettyprogress.Plan{
			{
				Bullet: prettyprogress.Running,
				Name:   "hello",
			},
			{
				Bullet: prettyprogress.Future,
				Name:   "",
			},
		}.String(),
		prettyprogress.Plan{
			{
				Bullet: prettyprogress.Running,
				Name:   "hello",
			},
			{
				Bullet: prettyprogress.Complete,
				Name:   "bye",
			},
		}.String(),
		prettyprogress.Plan{
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
