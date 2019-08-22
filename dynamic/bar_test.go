package dynamic_test

import (
	"testing"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/dynamic"
	"gotest.tools/assert"
)

func TestDynamic(t *testing.T) {
	var recieved []string
	updater := dynamic.NewProgressUpdater(100, 20, func(b string) {
		recieved = append(recieved, b)
	})

	ch := make(chan struct{})
	go func() {
		updater.Update(10)
		updater.Update(5)
		updater.Update(15)
		close(ch)
	}()

	<-ch

	assert.DeepEqual(t, recieved, []string{
		prettyprogress.NewBarWithWidth(10, 100, 20).String(),
		prettyprogress.NewBarWithWidth(5, 100, 20).String(),
		prettyprogress.NewBarWithWidth(15, 100, 20).String(),
	})
}
