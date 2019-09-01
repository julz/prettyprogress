package ui_test

import (
	"strings"
	"testing"

	"github.com/julz/prettyprogress/ui"
	"gotest.tools/assert"
)

func TestSteps(t *testing.T) {
	examples := []struct {
		Title  string
		Steps  ui.Steps
		Expect string
	}{
		{
			Title: "Basic",
			Steps: ui.Steps{
				{
					Name:   "Building..",
					Bullet: ui.Complete,
				},
				{
					Name:   "Downloading..",
					Bullet: ui.Downloading,
					Bar:    "[███         ]",
				},
				{
					Name:   "Scanning..",
					Bullet: ui.Running,
				},
				{
					Name:   "Waiting..",
					Bullet: ui.Future,
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
			Steps: ui.Steps{
				{
					Name:   "Building..",
					Bullet: ui.Complete,
				},
				{
					Name:   "Downloading..",
					Bullet: ui.Downloading,
					Bar:    "[███         ]",
				},
				{
					Name:   "This line is really really long..",
					Bullet: ui.Running,
				},
				{
					Name:   "Waiting..",
					Bullet: ui.Future,
				},
			},
			Expect: withoutPadding(`
				 ✓  Building..
				 ↡  Downloading..                       [███         ]
				 ►  This line is really really long..
				    Waiting..
			`),
		},
		{
			Title: "Colored bullets",
			Steps: ui.Steps{
				{
					Name:   "Building..",
					Bullet: ui.Complete,
					BulletColorFunc: func(s ...interface{}) string {
						assert.DeepEqual(t, s, []interface{}{
							"✓",
						})

						return "C"
					},
				},
			},
			Expect: withoutPadding(`
				 C  Building..
			`),
		},
	}

	for _, eg := range examples {
		t.Run(eg.Title, func(t *testing.T) {
			assert.Equal(t, eg.Expect, eg.Steps.String())
		})
	}
}

func withoutPadding(s string) string {
	result := ""
	for _, line := range strings.Split(s, "\n") {
		result += strings.TrimLeft(line, "\t") + "\n"
	}

	return strings.TrimRight(result, "\n") + "\n"
}
