package dynamic

import "github.com/julz/prettyprogress"

type StepUpdater struct {
	barWidth int
	barTotal int

	watcher stepWatcher
}

type stepWatcher func(s prettyprogress.Step)

type StatusUpdater interface {
	UpdateStatus(bullet prettyprogress.Bullet, status string)
	UpdateProgress(bullet prettyprogress.Bullet, status string, progress int)
}

func NewStatusUpdater(barTotal, barWidth int, w Watcher) *StepUpdater {
	return &StepUpdater{
		barWidth: barWidth,
		barTotal: barTotal,
		watcher:  func(s prettyprogress.Step) { w(s.String()) },
	}
}

func (b *StepUpdater) UpdateStatus(bullet prettyprogress.Bullet, status string) {
	b.update(bullet, status, "")
}

func (b *StepUpdater) Bar(bullet prettyprogress.Bullet, status string) *BarUpdater {
	return NewProgressUpdater(
		b.barTotal,
		b.barWidth,
		func(s string) {
			b.update(bullet, status, s)
		},
	)
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
	})
}
