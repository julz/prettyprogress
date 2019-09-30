package ui_test

import (
	"testing"

	"strings"

	"github.com/julz/prettyprogress/ui"
	"gotest.tools/assert"
)

func TestBar(t *testing.T) {
	examples := []struct {
		Title             string
		Bar               ui.Bar
		Expect            string
		ExpectLastBarChar string
	}{
		{
			Title:  "Simple empty bar",
			Bar:    ui.Bar{Progress: 0, Total: 4, Width: 4},
			Expect: "[    ]",
		},
		{
			Title:  "Empty bar larger width than total",
			Expect: "[        ]",
			Bar:    ui.Bar{Progress: 0, Total: 4, Width: 8},
		},
		{
			Title:  "Bar with progress",
			Expect: "[█   ]",
			Bar:    ui.Bar{Progress: 1, Total: 4, Width: 4},
		},
		{
			Title:  "Bar with progress, larger width than total",
			Expect: "[██      ]",
			Bar:    ui.Bar{Progress: 1, Total: 4, Width: 8},
		},
		{
			Title:  "Bar with 50% progress",
			Expect: "[████    ]",
			Bar:    ui.Bar{Progress: 2, Total: 4, Width: 8},
		},
		{
			Title:  "Bar with 50% progress, showing Percentage label",
			Expect: "[████    ] 50%",
			Bar:    ui.Bar{Progress: 2, Total: 4, Width: 8, LabelFunc: ui.PercentageLabel},
		},
		{
			Title:  "Non-integer progress step",
			Expect: "[███   ]",
			Bar:    ui.Bar{Progress: 2, Total: 4, Width: 6},
		},
		{
			Title:  "Unicode 1/8th characters when needed",
			Expect: "[████▏   ]",
			Bar:    ui.Bar{Progress: 33, Total: 64, Width: 8},
		},
		{
			Title:  "Unicode 2/8th characters when needed",
			Expect: "[████▎   ]",
			Bar:    ui.Bar{Progress: 34, Total: 64, Width: 8},
		},
		{
			Title:  "Unicode half-bar characters when needed",
			Expect: "[██▋  ]",
			Bar:    ui.Bar{Progress: 27, Total: 50, Width: 5},
		},
		{
			Title:  "Unicode 7/8th characters when needed",
			Expect: "[████▉   ]",
			Bar:    ui.Bar{Progress: 39, Total: 64, Width: 8},
		},
		{
			Title:  "Using full constructor",
			Expect: "[██▋  ]",
			Bar:    ui.NewBarWithWidth(27, 50, 5),
		},
		{
			Title:  "Using convenience constructor with default width",
			Expect: "[██████████          ]",
			Bar:    ui.NewBar(2, 4),
		},
		{
			Title:  "Full bar",
			Expect: "[█]",
			Bar:    ui.Bar{Progress: 1, Total: 1, Width: 1},
		},
		{
			Title:  "Over-full bar",
			Expect: "[█]",
			Bar:    ui.Bar{Progress: 2, Total: 1, Width: 1},
		},
		{
			Title:  "Bar with custom end chars",
			Expect: "┃██   ┃",
			Bar:    ui.Bar{Progress: 2, Total: 5, Width: 5, StartChar: "┃", EndChar: "┃"},
		},
	}

	for _, eg := range examples {
		t.Run(eg.Title, func(t *testing.T) {
			assert.Equal(t,
				len(eg.Expect),
				len(eg.Bar.String()),
				"incorrect length")
			assert.Equal(t,
				strings.Trim(eg.Expect, "[]█ "),
				strings.Trim(eg.Bar.String(), "[]█ "),
				"incorrect fractional characters")
			assert.Equal(t,
				strings.Trim(eg.Expect, "[] "),
				strings.Trim(eg.Bar.String(), "[] "),
				"incorrect filled section")
			assert.Equal(t, eg.Expect, eg.Bar.String())
		})
	}
}
