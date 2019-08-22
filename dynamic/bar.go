package dynamic

import (
	"github.com/julz/prettyprogress"
)

type Watcher func(s string)

type ProgressUpdater interface {
	UpdateProgress(progress int)
}

type BarUpdater struct {
	total int
	width int

	watcher Watcher
}

func NewProgressUpdater(total, width int, w Watcher) *BarUpdater {
	return &BarUpdater{
		total:   total,
		width:   width,
		watcher: w,
	}
}

func (b *BarUpdater) UpdateProgress(progress int) {
	b.watcher(prettyprogress.Bar{Progress: progress, Width: b.width, Total: b.total}.String())
}
