package updater

import (
	"github.com/julz/prettyprogress"
)

type Watcher func(s string)

type ProgressUpdater interface {
	UpdateProgress(progress int)
}

type Bar struct {
	total int
	width int

	watcher Watcher
}

func NewBar(total, width int, w Watcher) *Bar {
	return &Bar{
		total:   total,
		width:   width,
		watcher: w,
	}
}

func (b *Bar) UpdateProgress(progress int) {
	b.watcher(prettyprogress.Bar{Progress: progress, Width: b.width, Total: b.total}.String())
}
