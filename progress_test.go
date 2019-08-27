package prettyprogress_test

import (
	"strings"
	"testing"

	"github.com/julz/prettyprogress"
	"gotest.tools/assert"
)

func TestProgress(t *testing.T) {
	examples := []struct {
		Title  string
		Steps  prettyprogress.Steps
		Expect string
	}{
		{
			Title: "Basic",
			Steps: prettyprogress.Steps{
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
			},
			Expect: withoutPadding(`
				 ✓  Building..
				 ↡  Downloading..   [███         ]
				 ►  Scanning..
				    Waiting..
			`),
		},
		{
			Title: "Over-long label moves progress bar out",
			Steps: prettyprogress.Steps{
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
					Name:   "This line is really really long..",
					Bullet: prettyprogress.Running,
				},
				{
					Name:   "Waiting..",
					Bullet: prettyprogress.Future,
				},
			},
			Expect: withoutPadding(`
				 ✓  Building..
				 ↡  Downloading..                       [███         ]
				 ►  This line is really really long..
				    Waiting..
			`),
		},
	}

	for _, eg := range examples {
		t.Run(eg.Title, func(t *testing.T) {
			assert.Equal(t, eg.Expect, eg.Steps.String())
		})
	}
}

type StringableString string

func (s StringableString) String() string {
	return string(s)
}

func withoutPadding(s string) string {
	result := ""
	for _, line := range strings.Split(s, "\n") {
		result += strings.TrimLeft(line, "\t") + "\n"
	}

	return strings.TrimRight(result, "\n") + "\n"
}
