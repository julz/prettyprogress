package prettyprogress_test

import (
	"testing"

	"strings"

	"github.com/julz/prettyprogress"
	"gotest.tools/assert"
)

func TestBar(t *testing.T) {
	examples := []struct {
		Title             string
		Bar               prettyprogress.Bar
		Expect            string
		ExpectLastBarChar string
	}{
		{
			Title:  "Simple empty bar",
			Bar:    prettyprogress.Bar{Progress: 0, Total: 4, Width: 4},
			Expect: "[    ]",
		},
		{
			Title:  "Empty bar larger width than total",
			Expect: "[        ]",
			Bar:    prettyprogress.Bar{Progress: 0, Total: 4, Width: 8},
		},
		{
			Title:  "Bar with progress",
			Expect: "[█   ]",
			Bar:    prettyprogress.Bar{Progress: 1, Total: 4, Width: 4},
		},
		{
			Title:  "Bar with progress, larger width than total",
			Expect: "[██      ]",
			Bar:    prettyprogress.Bar{Progress: 1, Total: 4, Width: 8},
		},
		{
			Title:  "Bar with 50% progress",
			Expect: "[████    ]",
			Bar:    prettyprogress.Bar{Progress: 2, Total: 4, Width: 8},
		},
		{
			Title:  "Non-integer progress step",
			Expect: "[███   ]",
			Bar:    prettyprogress.Bar{Progress: 2, Total: 4, Width: 6},
		},
		{
			Title:  "Unicode half-bar characters when needed",
			Expect: "[██▌  ]",
			Bar:    prettyprogress.Bar{Progress: 2, Total: 4, Width: 5},
		},
		{
			Title:  "Using full constructor",
			Expect: "[██▌  ]",
			Bar:    prettyprogress.NewBarWithWidth(2, 4, 5),
		},
		{
			Title:  "Using convenience constructor with default width",
			Expect: "[██████████          ]",
			Bar:    prettyprogress.NewBar(2, 4),
		},
	}

	for _, eg := range examples {
		t.Run(eg.Title, func(t *testing.T) {
			assert.Equal(t, len(eg.Expect), len(eg.Bar.String()), "correct length")
			assert.Equal(t, strings.Trim(eg.Expect, "[] "), strings.Trim(eg.Bar.String(), "[] "), "correct filled section")
			assert.Equal(t, eg.Expect, eg.Bar.String())
		})
	}
}
