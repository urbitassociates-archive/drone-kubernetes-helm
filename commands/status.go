package commands

import (
	"os/exec"
)

func (c *Command) Status() error {
	cmd := exec.Command("helm", "status", c.Release)
	c.appendFlags(cmd)

	return run(cmd)
}
