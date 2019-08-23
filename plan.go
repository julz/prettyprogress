package prettyprogress

import (
	"fmt"
)

// Plan is a series of steps neccesary to complete the task
type Plan []Step

// Bullet is a unicode status icon for each Step in a Plan
type Bullet string

const (
	Failed      Bullet = "⚠"
	Future      Bullet = " "
	Running     Bullet = "►"
	Downloading Bullet = "↡"
	Uploading   Bullet = "↟"
	Complete    Bullet = "✓"
)

func (b Bullet) String() string {
	return string(b)
}

// Step represents a single step of the Plan
type Step struct {
	Bullet fmt.Stringer
	Name   string
	Bar    string
}

func (s Step) String() string {
	return fmt.Sprintf(" %s  %s   %s", s.Bullet, s.Name, s.Bar)
}

func (p Plan) String() string {
	s := "\n"
	for _, step := range p {
		s += step.String() + "\n"
	}

	return s
}
