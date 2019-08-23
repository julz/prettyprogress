package prettyprogress_test

import (
	"fmt"
	"testing"

	"github.com/julz/prettyprogress"
	"gotest.tools/assert"
)

func TestProgress(t *testing.T) {
	a := prettyprogress.Steps{
		{
			Name:   "Building..",
			Bullet: prettyprogress.Complete,
		},
		{
			Name:   "Downloading..",
			Bullet: prettyprogress.Downloading,
			Bar:    "[███         ]",
		},
		{
			Name:   "Scanning..",
			Bullet: prettyprogress.Running,
		},
		{
			Name:   "Waiting..",
			Bullet: prettyprogress.Future,
		},
	}

	assert.Equal(t, fmt.Sprintf("%s", a), `
 ✓  Building..
 ↡  Downloading..   [███         ]
 ►  Scanning..
    Waiting..
`)
}

type StringableString string

func (s StringableString) String() string {
	return string(s)
}
