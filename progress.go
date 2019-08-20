package prettyprogress

import (
	"fmt"
)

type Plan []Step

type Stringer interface {
	String() string
}

type Bullet string

const (
	Future      Bullet = " "
	Running     Bullet = "►"
	Downloading Bullet = "↡"
	Uploading   Bullet = "↟"
	Complete    Bullet = "✓"
)

func (b Bullet) String() string {
	return string(b)
}

type Step struct {
	Bullet fmt.Stringer
	Name   string
	Bar    fmt.Stringer
}

func (s Step) String() string {
	return fmt.Sprintf(" %s  %s   %s", s.Bullet, s.Name, emptyIfNil(s.Bar))
}

func (p Plan) String() string {
	s := "\n"
	for _, step := range p {
		s += step.String() + "\n"
	}

	return s
}

func emptyIfNil(s fmt.Stringer) string {
	if s == nil {
		return ""
	}

	return s.String()
}
