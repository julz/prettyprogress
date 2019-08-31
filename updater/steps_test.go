package updater_test

import (
	"fmt"
	"testing"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/updater"
	"gotest.tools/assert"
)

func TestMultistep(t *testing.T) {
	var recieved []string
	steps := updater.NewMultistep(
		func(s string) {
			recieved = append(recieved, s)
		},
		updater.WithBarWidth(20),
	)

	step1 := steps.AddStep(updater.WithBarTotal(100))
	step2 := steps.AddStep()

	step1.Update(prettyprogress.Running, "hello")
	step2.Update(prettyprogress.Complete, "bye")
	step1.UpdateWithProgress(prettyprogress.Downloading, "updated", 12)

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

func TestMultistepWithColor(t *testing.T) {
	var recieved []string
	steps := updater.NewMultistep(
		func(s string) {
			recieved = append(recieved, s)
		},
		updater.WithBarWidth(20),
		updater.WithBulletColor(
			prettyprogress.Downloading,
			func(s ...interface{}) string {
				return "ORANGE<" + fmt.Sprint(s...) + ">"
			},
		),
		updater.WithBulletColor(
			prettyprogress.Complete,
			func(s ...interface{}) string {
				return "GREEN<" + fmt.Sprint(s...) + ">"
			},
		),
	)

	step1 := steps.AddStep(updater.WithBarTotal(100))
	step2 := steps.AddStep()

	step1.Update(prettyprogress.Running, "hello")
	step2.Update(prettyprogress.Complete, "bye")
	step1.UpdateWithProgress(prettyprogress.Downloading, "updated", 12)

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
				Bullet: "GREEN<" + prettyprogress.Complete + ">",
				Name:   "bye",
			},
		}.String(),
		prettyprogress.Steps{
			{
				Bullet: "ORANGE<" + prettyprogress.Downloading + ">",
				Name:   "updated",
				Bar:    prettyprogress.NewBarWithWidth(12, 100, 20).String(),
			},
			{
				Bullet: "GREEN<" + prettyprogress.Complete + ">",
				Name:   "bye",
			},
		}.String(),
	})
}
