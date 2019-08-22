package dynamic

import (
	"sync"

	"github.com/julz/prettyprogress"
)

type Bar struct {
	mu  sync.RWMutex
	bar prettyprogress.Bar

	wmu     sync.RWMutex
	watcher func(prettyprogress.Bar)
}

func NewBar(total, width int) *Bar {
	return &Bar{
		bar:     prettyprogress.NewBarWithWidth(0, total, width),
		watcher: func(prettyprogress.Bar) {},
	}
}

func (b *Bar) Set(progress int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.bar.Progress = progress

	b.wmu.RLock()
	defer b.wmu.RUnlock()
	b.watcher(b.bar)
}

func (b *Bar) Watch(fn func(prettyprogress.Bar)) {
	b.wmu.Lock()
	defer b.wmu.Unlock()

	b.watcher = fn
}

func (b *Bar) String() string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.bar.String()
}
