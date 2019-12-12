package prettyprogress

import (
	"github.com/julz/prettyprogress/ui"
)

// Step is an updater whose watcher recieves the output of
// ui.Step's String() method whenever its UpdateStatus method is
// called
type Step struct {
	barWidth int
	barTotal int
	barLabel ui.LabelFunc

	bullets ui.BulletSet
	colors  Colors

	watcher stepWatcher
}

type stepWatcher func(s ui.Step)

// NewStep creates a new Step which can be updated with the status of a single task
func NewStep(barTotal, barWidth int, print PrintFunc) *Step {
	return &Step{
		barWidth: barWidth,
		barTotal: barTotal,
		bullets:  ui.DefaultBulletSet,
		colors:   DefaultColors,
		watcher:  func(s ui.Step) { print(s.String()) },
	}
}

// Fail sets the steps name to the given string and changes the bullet to a symbol indicating failure
func (b *Step) Fail(msg string) {
	b.Update(b.bullets.Failed, msg)
}

// Complete sets the steps name to the given string and changes the bullet to a
// symbol indicating the task has been completed
func (b *Step) Complete(msg string) {
	b.Update(b.bullets.Complete, b.colors.Completed(msg))
}

// Start sets the steps name to the given string and changes the bullet to a symbol indicating the task is running
func (b *Step) Start(msg string) {
	b.Update(b.bullets.Running, msg)
}

// UpdateState updates the step with the given status and the bullet from the current BulletSet corresponding to the given state
func (b *Step) UpdateState(state ui.BulletState, status string) {
	switch state {
	case ui.RunningState:
		b.Start(status)
	case ui.CompleteState:
		b.Complete(status)
	case ui.FailedState:
		b.Fail(status)
	default:
		b.update(ui.Future, status, "")
	}
}

// Update sets the steps name to the givem status and updated the bullet to the given Bullet
func (b *Step) Update(bullet ui.Bullet, status string) {
	b.update(bullet, status, "")
}

// Bar returns just the Bar component of the Step which can be updated to show numeric
// progress of the current task
//
// For example, this could be used with a function that expects an interface with
// an Update method as follows:
//
//   download(step.Bar(ui.Downloading, "Downloading layer.."))
//
// The corresponding `download` method could look as follows:
//
//   bar(p type interface{ Update(int) }) {
//     p.Update(10)
//     p.Update(30)
//     p.Update(100)
//   })
func (b *Step) Bar(bullet ui.Bullet, status string) *Bar {
	return &Bar{
		total: b.barTotal,
		width: b.barWidth,
		label: b.barLabel,
		printFunc: func(s string) {
			b.update(bullet, status, s)
		},
	}
}

// UpdateWithProgress updates the Bullet, Status and Progress Bar of the current
// step. Often either one of the convenience methods like Start,
// Fail, Complete, or Bar will be a better option.
func (b *Step) UpdateWithProgress(bullet ui.Bullet, status string, progress int) {
	b.update(bullet, status, ui.Bar{
		Width:    b.barWidth,
		Total:    b.barTotal,
		Progress: progress,
	}.String())
}

func (b *Step) update(bullet ui.Bullet, status, bar string) {
	b.watcher(ui.Step{
		Bullet: bullet,
		Name:   status,
		Bar:    bar,
	})
}
