package prettyprogress_test

import (
	"fmt"
	"testing"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/ui"
	"gotest.tools/assert"
)

func TestMultistep(t *testing.T) {
	var recieved []string
	steps := prettyprogress.NewMultistep(
		func(s string) {
			recieved = append(recieved, s)
		},
		prettyprogress.WithBarWidth(20),
	)

	step1 := steps.AddStep(prettyprogress.WithBarTotal(100))
	step2 := steps.AddStep()

	step1.Update(ui.Running, "hello")
	step2.Update(ui.Complete, "bye")
	step1.UpdateWithProgress(ui.Downloading, "updated", 12)

	assert.DeepEqual(t, recieved, []string{
		ui.Steps{
			{
				Bullet: ui.Running,
				Name:   "hello",
			},
			{
				Bullet: ui.Future,
				Name:   "",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Running,
				Name:   "hello",
			},
			{
				Bullet: ui.Complete,
				Name:   "bye",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Downloading,
				Name:   "updated",
				Bar:    ui.NewBarWithWidth(12, 100, 20).String(),
			},
			{
				Bullet: ui.Complete,
				Name:   "bye",
			},
		}.String(),
	})
}

func TestMultistepWithColor(t *testing.T) {
	var recieved []string
	steps := prettyprogress.NewMultistep(
		func(s string) {
			recieved = append(recieved, s)
		},
		prettyprogress.WithBarWidth(20),
		prettyprogress.WithBulletColor(
			ui.Downloading,
			func(s ...interface{}) string {
				return "ORANGE<" + fmt.Sprint(s...) + ">"
			},
		),
		prettyprogress.WithBulletColor(
			ui.Complete,
			func(s ...interface{}) string {
				return "GREEN<" + fmt.Sprint(s...) + ">"
			},
		),
	)

	step1 := steps.AddStep(prettyprogress.WithBarTotal(100))
	step2 := steps.AddStep()

	step1.Update(ui.Running, "hello")
	step2.Update(ui.Complete, "bye")
	step1.UpdateWithProgress(ui.Downloading, "updated", 12)

	assert.DeepEqual(t, recieved, []string{
		ui.Steps{
			{
				Bullet: ui.Running,
				Name:   "hello",
			},
			{
				Bullet: ui.Future,
				Name:   "",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Running,
				Name:   "hello",
			},
			{
				Bullet: "GREEN<" + ui.Complete + ">",
				Name:   "bye",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: "ORANGE<" + ui.Downloading + ">",
				Name:   "updated",
				Bar:    ui.NewBarWithWidth(12, 100, 20).String(),
			},
			{
				Bullet: "GREEN<" + ui.Complete + ">",
				Name:   "bye",
			},
		}.String(),
	})
}
