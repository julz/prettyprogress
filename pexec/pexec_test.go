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
		UpdateFunc: func(bullet ui.BulletState, status string) {
			updates = append(updates, updateArgs{
				Bullet: bullet,
				Status: status,
			})
		},
	})

	assert.DeepEqual(t, updates, []updateArgs{
		{
			Status: "Waiting to Run 'echo hello'",
		},
	})

	cmd.Run()

	assert.Error(t, echo.Start(), "exec: already started")

	assert.DeepEqual(t, updates, []updateArgs{
		{
			Status: "Waiting to Run 'echo hello'",
		},
		{
			Bullet: ui.RunningState,
			Status: "Running 'echo hello'..",
		},
		{
			Bullet: ui.CompleteState,
			Status: "Finished 'echo hello'",
		},
	})
}

type updateArgs struct {
	Bullet ui.BulletState
	Status string
}

type updater struct {
	UpdateFunc func(bullet ui.BulletState, status string)
}

func (u updater) UpdateState(bullet ui.BulletState, status string) {
	u.UpdateFunc(bullet, status)
}
