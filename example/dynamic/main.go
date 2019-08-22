package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/prettyprogress"
	"github.com/julz/prettyprogress/dynamic"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	bar := dynamic.NewBar(100, 20)
	bar.Watch(func(b prettyprogress.Bar) {
		fmt.Fprintln(w, b.String())
		w.Flush()
	})

	doSomethingWithProgress(bar)
}

func doSomethingWithProgress(b *dynamic.Bar) {
	for i := 0; i <= 100; i++ {
		b.Set(i)
		time.Sleep(5 * time.Millisecond)
	}
}
