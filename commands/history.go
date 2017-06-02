package commands

import (
	"os/exec"
)

func (c *Command) History() error {
	cmd := exec.Command("helm", "history", c.Release)
	c.appendFlags(cmd)

	return run(cmd)
}
