package updater

import (
	"sync"

	"github.com/julz/prettyprogress"
)

type Plan struct {
	watcher Watcher

	barWidth int

	mu    sync.Mutex
	steps prettyprogress.Plan
}

func NewMultistep(barWidth int, watcher Watcher) *Plan {
	return &Plan{
		barWidth: barWidth,
		watcher:  watcher,
	}
}

func (p *Plan) AddStep(total int) *Step {
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
