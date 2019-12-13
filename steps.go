package prettyprogress

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/julz/prettyprogress/ui"
)

const defaultBarTotal = 100

type Steps struct {
	printFunc PrintFunc

	barWidth int
	barLabel ui.LabelFunc
	bullets  ui.BulletSet
	colors   Colors

	frame         int
	frameMu       sync.RWMutex
	frameTickerCh <-chan time.Time

	mu    sync.RWMutex
	steps ui.Steps
}

// PrintFunc is a function that is called when new versions of a Bar, Step or Multistep
// are created as a result of calling UpdateX methods.
type PrintFunc func(s string)

type flusher interface {
	Flush() error
}

// NewMultistep creates a new updater that can have multiple sub-steps. When any of the steps
// are updated, the Writer is called with the new string to display.
func NewMultistep(writer io.Writer, options ...StepsOption) *Steps {
	s := &Steps{
		barWidth: 20,
		bullets:  ui.DefaultBulletSet,
		printFunc: func(w string) {
			writer.Write([]byte(w))
			if writer, ok := writer.(flusher); ok {
				writer.Flush()
			}
		},
		colors: DefaultColors,
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

// NewFancyMultistep creates a new updater that prints step progress with fancy
// colours, bullets and animations using the given Watcher
func NewFancyMultistep(writer io.Writer, extraOptions ...StepsOption) *Steps {
	return NewMultistep(writer,
		WithBullets(ui.ColoredAnimatedBulletSet),
		WithAnimationFrameTicker(time.NewTicker(200*time.Millisecond).C),
		WithLabelColors(Colors{
			Future:    color.New(color.FgHiBlack).Sprint,
			Completed: color.New(color.FgHiBlack).Sprint,
		}))
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
		p.steps[stepIndex] = s
		p.mu.Unlock()

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

	p.mu.RLock()
	defer p.mu.RUnlock()

	p.printFunc(p.steps.AnimatedString(p.frame))
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

func WithLabelColors(colors Colors) func(*Steps) {
	return func(s *Steps) {
		s.colors = colors
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
