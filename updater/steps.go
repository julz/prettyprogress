package updater

import (
	"sync"

	"github.com/julz/prettyprogress"
)

type Steps struct {
	watcher Watcher

	barWidth     int
	bulletColors map[string]func(a ...interface{}) string

	mu    sync.Mutex
	steps prettyprogress.Steps
}

func NewMultistep(watcher Watcher, options ...StepsOption) *Steps {
	s := &Steps{
		bulletColors: make(map[string]func(a ...interface{}) string),
		barWidth:     20,
		watcher:      watcher,
	}

	for _, option := range options {
		option(s)
	}

	return s
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
		bulletColors: p.bulletColors,

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

type StepsOption func(s *Steps)

func WithBulletColor(bullet prettyprogress.Bullet, color func(...interface{}) string) func(s *Steps) {
	return func(s *Steps) {
		s.bulletColors[bullet.String()] = color
	}
}

func WithBarWidth(width int) func(s *Steps) {
	return func(s *Steps) {
		s.barWidth = width
	}
}
