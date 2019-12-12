package pexec

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/julz/prettyprogress/ui"
)

type Updater interface {
	UpdateState(state ui.BulletState, status string)
}

type Cmd struct {
	cmd     *exec.Cmd
	name    string
	updater Updater
}

func Wrap(cmd *exec.Cmd, u Updater) *Cmd {
	return Wrapn(strings.Join(cmd.Args, " "), cmd, u)
}

func Wrapn(name string, cmd *exec.Cmd, u Updater) *Cmd {
	c := &Cmd{
		cmd:     cmd,
		updater: u,
		name:    name,
	}

	u.UpdateState("", fmt.Sprintf("Waiting to Run '%s'", name))
	return c
}

func (c *Cmd) Run() error {
	c.updater.UpdateState(ui.RunningState, fmt.Sprintf("Running '%s'..", c.name))
	err := c.cmd.Run()
	c.updater.UpdateState(ui.CompleteState, fmt.Sprintf("Finished '%s'", c.name))

	return err
}
