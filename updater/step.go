package updater

import "github.com/julz/prettyprogress"

// Step is an updater whose watcher recieves the output of
// prettyprogress.Step's String() method whenever its UpdateStatus method is
// called
type Step struct {
	bulletColors map[string]func(a ...interface{}) string
	barWidth     int
	barTotal     int

	watcher stepWatcher
}

type stepWatcher func(s prettyprogress.Step)

// NewStep creates a new Step which can be updated with the status of a single task
func NewStep(barTotal, barWidth int, w Watcher) *Step {
	return &Step{
		barWidth: barWidth,
		barTotal: barTotal,
		watcher:  func(s prettyprogress.Step) { w(s.String()) },
	}
}

// Fail sets the steps name to the given string and changes the bullet to a symbol indicating failure
func (b *Step) Fail(msg string) {
	b.UpdateStatus(prettyprogress.Failed, msg)
}

// Complete sets the steps name to the given string and changes the bullet to a
// symbol indicating the task has been completed
func (b *Step) Complete(msg string) {
	b.UpdateStatus(prettyprogress.Complete, msg)
}

// Start sets the steps name to the given string and changes the bullet to a symbol indicating the task is running
func (b *Step) Start(msg string) {
	b.UpdateStatus(prettyprogress.Running, msg)
}

// UpdateStatus sets the steps name to the givem status and updated the bullet to the given Bullet
func (b *Step) UpdateStatus(bullet prettyprogress.Bullet, status string) {
	b.update(bullet, status, "")
}

// Bar returns just the Bar component of the Step which can be updated to show numeric
// progress of the current task
//
// For example, this could be used with a download() function that takes an interface with an UpdateProgress method
// as follows:
//
//  download(bar step.Bar(prettyprogress.Downloading, "Download X..")) {
//   bar.UpdateProgress(10)
//   bar.UpdateProgress(100)
//  }
func (b *Step) Bar(bullet prettyprogress.Bullet, status string) *Bar {
	return NewBar(
		b.barTotal,
		b.barWidth,
		func(s string) {
			b.update(bullet, status, s)
		},
	)
}

// UpdateProgress updates the Bullet, Status and Progress Bar of the current
// step. Often either UpdateStatus, one of the convenience methods like Start,
// Fail, Complete, or Bar will be a better option.
func (b *Step) UpdateProgress(bullet prettyprogress.Bullet, status string, progress int) {
	b.update(bullet, status, prettyprogress.Bar{
		Width:    b.barWidth,
		Total:    b.barTotal,
		Progress: progress,
	}.String())
}

func (b *Step) update(bullet prettyprogress.Bullet, status, bar string) {
	b.watcher(prettyprogress.Step{
		Bullet:          bullet,
		BulletColorFunc: b.bulletColors[bullet.String()],
		Name:            status,
		Bar:             bar,
	})
}
