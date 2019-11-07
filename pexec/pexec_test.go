package pexec_test

import (
	"os/exec"
	"testing"

	"github.com/julz/prettyprogress/pexec"
	"github.com/julz/prettyprogress/ui"
	"gotest.tools/assert"
)

func TestPexec(t *testing.T) {
	var updates []updateArgs
	echo := exec.Command("echo", "hello")
	cmd := pexec.Wrap(echo, updater{
		UpdateFunc: func(bullet ui.Bullet, status string) {
			updates = append(updates, updateArgs{
				Bullet: bullet,
				Status: status,
			})
		},
	})

	assert.DeepEqual(t, updates, []updateArgs{
		{
			Bullet: ui.Future,
			Status: "Run 'echo hello'",
		},
	})

	cmd.Run()

	assert.Error(t, echo.Start(), "exec: already started")

	assert.DeepEqual(t, updates, []updateArgs{
		{
			Bullet: ui.Future,
			Status: "Run 'echo hello'",
		},
		{
			Bullet: ui.Running,
			Status: "Running 'echo hello'..",
		},
		{
			Bullet: ui.Complete,
			Status: "Finished 'echo hello'",
		},
	})
}

type updateArgs struct {
	Bullet ui.Bullet
	Status string
}

type updater struct {
	UpdateFunc func(bullet ui.Bullet, status string)
}

func (u updater) Update(bullet ui.Bullet, status string) {
	u.UpdateFunc(bullet, status)
}
