package dynamic_test

import (
	"testing"

	"time"

	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/dynamic"
	"gotest.tools/assert"
)

func TestDynamic(t *testing.T) {
	t.Run("behaves the same as the static version", func(t *testing.T) {
		bar := dynamic.NewBar(100, 20)
		bar.Set(5)

		assert.Equal(t,
			prettyprogress.NewBarWithWidth(5, 100, 20).String(),
			bar.String())
	})

	t.Run("is thread-safe", func(t *testing.T) {
		if !RaceOn {
			t.Skip("race detector not running so this test does nothing")
		}

		bar := dynamic.NewBar(100, 20)

		go func() {
			bar.Set(5)
		}()

		expect := prettyprogress.NewBarWithWidth(5, 100, 20).String()
		for i := 0; i < 100; i++ {
			if expect == bar.String() {
				return
			}

			time.Sleep(5 * time.Millisecond)
		}

		t.Fatalf("expected to eventually see the correct String()")
	})

	t.Run("can be watched for updates", func(t *testing.T) {
		bar := dynamic.NewBar(100, 20)

		ch := make(chan struct{})
		go func() {
			bar.Set(10)
			bar.Set(5)
			bar.Set(15)
			close(ch)
		}()

		var recieved []prettyprogress.Bar
		bar.Watch(func(b prettyprogress.Bar) {
			recieved = append(recieved, b)
		})

		<-ch

		assert.DeepEqual(t, recieved, []prettyprogress.Bar{
			prettyprogress.NewBarWithWidth(10, 100, 20),
			prettyprogress.NewBarWithWidth(5, 100, 20),
			prettyprogress.NewBarWithWidth(15, 100, 20),
		})
	})
}
