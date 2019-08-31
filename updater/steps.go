package updater

import (
	"sync"

	"github.com/julz/prettyprogress"
)

const defaultBarTotal = 100

type Steps struct {
	watcher Watcher

	barWidth     int
	bulletColors map[string]func(a ...interface{}) string

	mu    sync.Mutex
	steps prettyprogress.Steps
}

// NewMultistep creates a new updater that can have multiple sub-steps. When any of the steps
// are updated, the Watcher is called with the new string to display.
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

// AddStep adds a sub-step to the display
func (p *Steps) AddStep(options ...StepOption) *Step {
	p.mu.Lock()
	stepIndex := len(p.steps)
	p.steps = append(p.steps, prettyprogress.Step{Bullet: prettyprogress.Future})
	p.mu.Unlock()

	s := &Step{
		bulletColors: p.bulletColors,

		barWidth: p.barWidth,
		barTotal: defaultBarTotal,
		watcher: func(s prettyprogress.Step) {
			p.mu.Lock()
			defer p.mu.Unlock()

			p.steps[stepIndex] = s
			p.watcher(p.steps.String())
		},
	}

	for _, o := range options {
		o(s)
	}

	return s
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

type StepOption func(s *Step)

func WithBarTotal(total int) func(s *Step) {
	return func(s *Step) {
		s.barTotal = total
	}
}

func WithStatus(msg string) func(S *Step) {
	return func(s *Step) {
		s.Update(prettyprogress.Future, msg)
	}
}
