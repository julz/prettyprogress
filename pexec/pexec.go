package pexec

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/julz/prettyprogress/ui"
)

type Updater interface {
	Update(bullet ui.Bullet, status string)
}

type Cmd struct {
	cmd     *exec.Cmd
	updater Updater
}

func Wrap(cmd *exec.Cmd, u Updater) *Cmd {
	c := &Cmd{
		cmd:     cmd,
		updater: u,
	}

	u.Update(ui.Future, fmt.Sprintf("Run '%s'", strings.Join(cmd.Args, " ")))
	return c
}

func (c *Cmd) Run() error {
	c.updater.Update(ui.Running, fmt.Sprintf("Running '%s'..", strings.Join(c.cmd.Args, " ")))
	err := c.cmd.Run()
	c.updater.Update(ui.Complete, fmt.Sprintf("Finished '%s'", strings.Join(c.cmd.Args, " ")))

	return err
}
