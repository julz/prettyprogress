package prettyprogress_test

import (
	"testing"

	"github.com/julz/prettyprogress"
	"gotest.tools/assert"
)

func TestBar(t *testing.T) {
	assert.Equal(t, "[    ]", prettyprogress.Bar{Progress: 0, Total: 4, Width: 4}.String())
	assert.Equal(t, "[        ]", prettyprogress.Bar{Progress: 0, Total: 4, Width: 8}.String())
	assert.Equal(t, "[█   ]", prettyprogress.Bar{Progress: 1, Total: 4, Width: 4}.String())
	assert.Equal(t, "[██      ]", prettyprogress.Bar{Progress: 1, Total: 4, Width: 8}.String())
	assert.Equal(t, "[████    ]", prettyprogress.Bar{Progress: 2, Total: 4, Width: 8}.String())
	assert.Equal(t, "[███   ]", prettyprogress.Bar{Progress: 2, Total: 4, Width: 6}.String())
	assert.Equal(t, "[██▌  ]", prettyprogress.Bar{Progress: 2, Total: 4, Width: 5}.String(), "should use unicode half-bar characters")
}
