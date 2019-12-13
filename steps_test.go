package prettyprogress_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/ui"
	"gotest.tools/assert"
)

func TestMultistep(t *testing.T) {
	bullets := ui.DefaultBulletSet
	bullets.Running = ui.Bullet{"OVERRIDDEN"}

	var recieved []string
	steps := prettyprogress.NewMultistep(
		writer(func(s string) {
			recieved = append(recieved, s)
		}),
		prettyprogress.WithBarWidth(20),
		prettyprogress.WithBullets(bullets),
		prettyprogress.WithLabelColors(prettyprogress.Colors{
			Future:    func(s ...interface{}) string { return fmt.Sprintf("FUTURE<%s>", s...) },
			Completed: func(s ...interface{}) string { return fmt.Sprintf("COMPLETED<%s>", s...) },
		}),
	)

	step1 := steps.AddStep("", 100)
	step2 := steps.AddStep("future", 0)

	step1.Start("hello")
	step2.Complete("bye")
	step1.UpdateWithProgress(ui.Downloading, "updated", 12)

	assert.DeepEqual(t, recieved, []string{
		ui.Steps{
			{
				Bullet: ui.Future,
				Name:   "",
			},
			{
				Bullet: ui.Future,
				Name:   "FUTURE<future>",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"OVERRIDDEN"},
				Name:   "hello",
			},
			{
				Bullet: ui.Future,
				Name:   "FUTURE<future>",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"OVERRIDDEN"},
				Name:   "hello",
			},
			{
				Bullet: ui.Complete,
				Name:   "COMPLETED<bye>",
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
				Name:   "COMPLETED<bye>",
			},
		}.String(),
	})
}

func TestMultistepAnimations(t *testing.T) {
	bullets := ui.DefaultBulletSet
	bullets.Running = ui.Bullet{"1", "2", "3"}

	c := make(chan time.Time)
	defer close(c)

	var recieved []string
	steps := prettyprogress.NewMultistep(
		writer(func(s string) {
			recieved = append(recieved, s)
		}),
		prettyprogress.WithBarWidth(20),
		prettyprogress.WithBullets(bullets),
		prettyprogress.WithAnimationFrameTicker(c),
	)

	step1 := steps.AddStep("hello", 100)
	step1.Start("hello")

	assert.DeepEqual(t, recieved, []string{
		ui.Steps{
			{
				Bullet: ui.Bullet{" "},
				Name:   "hello",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"1"},
				Name:   "hello",
			},
		}.String(),
	})

	c <- time.Now()
	c <- time.Now()

	time.Sleep(100 * time.Millisecond) // :-(

	assert.DeepEqual(t, recieved, []string{
		ui.Steps{
			{
				Bullet: ui.Bullet{" "},
				Name:   "hello",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"1"},
				Name:   "hello",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"2"},
				Name:   "hello",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"3"},
				Name:   "hello",
			},
		}.String(),
	})

	c <- time.Now()
	time.Sleep(100 * time.Millisecond) // :-(

	assert.DeepEqual(t, recieved, []string{
		ui.Steps{
			{
				Bullet: ui.Bullet{" "},
				Name:   "hello",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"1"},
				Name:   "hello",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"2"},
				Name:   "hello",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"3"},
				Name:   "hello",
			},
		}.String(),
		ui.Steps{
			{
				Bullet: ui.Bullet{"1"}, // should wrap around
				Name:   "hello",
			},
		}.String(),
	})

}

type writer func(w string)

func (fn writer) Write(b []byte) (n int, e error) {
	fn(string(b))
	return 0, nil
}
