package prettyprogress

import (
	"fmt"
)

// Steps is a series of steps neccesary to complete the task, w
type Steps []Step

// String outputs the set of steps as a nicely formatted String
func (p Steps) String() string {
	s := "\n"
	for _, step := range p {
		s += step.String() + "\n"
	}

	return s
}

// Step represents a single step
type Step struct {
	Bullet fmt.Stringer
	Name   string
	Bar    string
}

// String outputs the Step as a nicely formatted String
func (s Step) String() string {
	return fmt.Sprintf(" %s  %s   %s", s.Bullet, s.Name, s.Bar)
}

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
