package prettyprogress

import (
	"fmt"
	"sync"
	"time"

	"github.com/julz/prettyprogress/ui"
)

const defaultBarTotal = 100

type Steps struct {
	watcher Watcher

	barWidth int
	barLabel ui.LabelFunc
	bullets  ui.BulletSet
	colors   Colors

	frame         int
	frameMu       sync.RWMutex
	frameTickerCh <-chan time.Time

	mu    sync.Mutex
	steps ui.Steps
}

// NewMultistep creates a new updater that can have multiple sub-steps. When any of the steps
// are updated, the Watcher is called with the new string to display.
func NewMultistep(watcher Watcher, options ...StepsOption) *Steps {
	s := &Steps{
		barWidth: 20,
		bullets:  ui.DefaultBulletSet,
		watcher:  watcher,
		colors:   DefaultColors,
	}

	for _, option := range options {
		option(s)
	}

	go func() {
		for range s.frameTickerCh {
			s.frameMu.Lock()
			s.frame = s.frame + 1
			s.frameMu.Unlock()

			s.refresh()
		}
	}()

	return s
}

// AddStep adds a sub-step to the display
func (p *Steps) AddStep(name string, barTotal int) *Step {
	p.mu.Lock()
	stepIndex := len(p.steps)
	p.steps = append(p.steps, ui.Step{Bullet: ui.Future})
	p.mu.Unlock()

	if barTotal == 0 {
		barTotal = 100
	}

	s := NewStep(barTotal, p.barWidth, nil)
	s.barLabel = p.barLabel
	s.bullets = p.bullets
	s.colors = p.colors
	s.watcher = func(s ui.Step) {
		p.mu.Lock()
		defer p.mu.Unlock()

		p.steps[stepIndex] = s
		p.refresh()
	}

	if name != "" {
		s.update(ui.Future, p.colors.Future(name), "")
	}

	return s
}

func (p *Steps) refresh() {
	p.frameMu.RLock()
	defer p.frameMu.RUnlock()

	p.watcher(p.steps.AnimatedString(p.frame))
}

type StepsOption func(s *Steps)

func WithBarWidth(width int) func(s *Steps) {
	return func(s *Steps) {
		s.barWidth = width
	}
}

func WithBarLabel(fn ui.LabelFunc) func(*Steps) {
	return func(s *Steps) {
		s.barLabel = fn
	}
}

func WithBullets(b ui.BulletSet) func(*Steps) {
	return func(s *Steps) {
		s.bullets = b
	}
}

func WithAnimationFrameTicker(c <-chan time.Time) func(*Steps) {
	return func(s *Steps) {
		s.frameTickerCh = c
	}
}

type Colors struct {
	Future    func(s ...interface{}) string
	Completed func(s ...interface{}) string
}

var DefaultColors = Colors{
	Future:    func(s ...interface{}) string { return fmt.Sprintf("%s", s...) },
	Completed: func(s ...interface{}) string { return fmt.Sprintf("%s", s...) },
}

func WithLabelColors(colors Colors) func(*Steps) {
	return func(s *Steps) {
		s.colors = colors
	}
}
