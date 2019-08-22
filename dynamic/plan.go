package dynamic

import (
	"sync"

	"github.com/julz/prettyprogress"
)

type PlanUpdater struct {
	watcher Watcher

	barWidth int

	mu    sync.Mutex
	steps prettyprogress.Plan
}

func NewMultistepUpdater(barWidth int, watcher Watcher) *PlanUpdater {
	return &PlanUpdater{
		barWidth: barWidth,
		watcher:  watcher,
	}
}

func (p *PlanUpdater) AddStep(total int) *StepUpdater {
	p.mu.Lock()
	stepIndex := len(p.steps)
	p.steps = append(p.steps, prettyprogress.Step{Bullet: prettyprogress.Future})
	p.mu.Unlock()

	return &StepUpdater{
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
