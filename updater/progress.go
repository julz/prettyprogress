package updater

import (
	"sync"

	"github.com/julz/prettyprogress"
)

type Steps struct {
	watcher Watcher

	barWidth int

	mu    sync.Mutex
	steps prettyprogress.Steps
}

func NewMultistep(barWidth int, watcher Watcher) *Steps {
	return &Steps{
		barWidth: barWidth,
		watcher:  watcher,
	}
}

func (p *Steps) AddStepWithStatus(status string, total int) *Step {
	s := p.AddStep(total)
	s.UpdateStatus(prettyprogress.Future, status)
	return s
}

func (p *Steps) AddStep(total int) *Step {
	p.mu.Lock()
	stepIndex := len(p.steps)
	p.steps = append(p.steps, prettyprogress.Step{Bullet: prettyprogress.Future})
	p.mu.Unlock()

	return &Step{
		barWidth: p.barWidth,
		barTotal: total,
		watcher: func(s prettyprogress.Step) {
			p.mu.Lock()
			defer p.mu.Unlock()

			p.steps[stepIndex] = s
			p.watcher(p.steps.String())
		},
	}
}
