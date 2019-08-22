package updater

import "github.com/julz/prettyprogress"

type Step struct {
	barWidth int
	barTotal int

	watcher stepWatcher
}

type stepWatcher func(s prettyprogress.Step)

type StatusUpdater interface {
	UpdateStatus(bullet prettyprogress.Bullet, status string)
	UpdateProgress(bullet prettyprogress.Bullet, status string, progress int)
}

func NewStep(barTotal, barWidth int, w Watcher) *Step {
	return &Step{
		barWidth: barWidth,
		barTotal: barTotal,
		watcher:  func(s prettyprogress.Step) { w(s.String()) },
	}
}

func (b *Step) UpdateStatus(bullet prettyprogress.Bullet, status string) {
	b.update(bullet, status, "")
}

func (b *Step) Bar(bullet prettyprogress.Bullet, status string) *Bar {
	return NewBar(
		b.barTotal,
		b.barWidth,
		func(s string) {
			b.update(bullet, status, s)
		},
	)
}

func (b *Step) UpdateProgress(bullet prettyprogress.Bullet, status string, progress int) {
	b.update(bullet, status, prettyprogress.Bar{
		Width:    b.barWidth,
		Total:    b.barTotal,
		Progress: progress,
	}.String())
}

func (b *Step) update(bullet prettyprogress.Bullet, status, bar string) {
	b.watcher(prettyprogress.Step{
		Bullet: bullet,
		Name:   status,
		Bar:    bar,
	})
}
