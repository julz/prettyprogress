package dynamic

import "github.com/julz/prettyprogress"

type StepUpdater struct {
	barWidth int
	barTotal int

	watcher Watcher
}

type StatusUpdater interface {
	Update(bullet prettyprogress.Bullet, status string)
	UpdateProgress(bullet prettyprogress.Bullet, status string, progress int)
}

func NewStatusUpdater(barTotal, barWidth int, w Watcher) *StepUpdater {
	return &StepUpdater{
		barWidth: barWidth,
		barTotal: barTotal,
		watcher:  w,
	}
}

func (b *StepUpdater) Update(bullet prettyprogress.Bullet, status string) {
	b.update(bullet, status, "")
}

func (b *StepUpdater) UpdateProgress(bullet prettyprogress.Bullet, status string, progress int) {
	b.update(bullet, status, prettyprogress.Bar{
		Width:    b.barWidth,
		Total:    b.barTotal,
		Progress: progress,
	}.String())
}

func (b *StepUpdater) update(bullet prettyprogress.Bullet, status, bar string) {
	b.watcher(prettyprogress.Step{
		Bullet: bullet,
		Name:   status,
		Bar:    bar,
	}.String())
}
