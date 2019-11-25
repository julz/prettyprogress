package ui

import (
	"fmt"
	"strings"
)

// Steps is a series of steps neccesary to complete the task
type Steps []Step

// String outputs the set of steps as a nicely formatted String
func (p Steps) String() string {
	return p.AnimatedString(0)
}

// AnimatedString outputs the steps as a nicely formatted string with the
// given frame of any bullets shown
func (p Steps) AnimatedString(frame int) string {
	longestName := 0
	for _, step := range p {
		if len(step.Name) > longestName {
			longestName = len(step.Name)
		}
	}

	s := "\n"
	for _, step := range p {
		step.paddedName = step.Name + strings.Repeat(" ", longestName-len(step.Name))
		s += step.AnimatedString(frame) + "\n"
	}

	return s
}

// Step represents a single step
// It is formatted as `$Bullet $Name $Bar` with appropriate spacing
type Step struct {
	Bullet              fmt.Stringer
	BulletColorFunc     func(...interface{}) string
	BulletAnimationFunc func(b string, frame int) string

	Name string

	Bar string

	paddedName string
}

// String outputs the Step as a nicely formatted String
func (s Step) String() string {
	return s.AnimatedString(0)
}

// AnimatedString returns the Step's string output for a given frame
func (s Step) AnimatedString(frame int) string {
	if s.paddedName == "" {
		s.paddedName = s.Name
	}

	bullet := s.Bullet.String()
	if s.BulletColorFunc != nil {
		bullet = s.BulletColorFunc(s.Bullet.String())
	}

	if s.BulletAnimationFunc != nil {
		bullet = s.BulletAnimationFunc(bullet, frame)
	}

	name := s.Name
	if s.Bar != "" {
		name = s.paddedName + "   "
	}

	return fmt.Sprintf(" %s  %s%s", bullet, name, s.Bar)
}

// Bullet is a unicode status icon for a Step
type Bullet string

const (
	Failed      Bullet = "✗"
	Future      Bullet = " "
	Running     Bullet = "►"
	Downloading Bullet = "↡"
	Uploading   Bullet = "↟"
	Complete    Bullet = "✓"
)

func (b Bullet) String() string {
	return string(b)
}
