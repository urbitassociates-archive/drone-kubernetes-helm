package commands

import (
	"os/exec"
)

func (c *Command) Test() error {
	cmd := exec.Command("helm", "test", c.Release)
	c.appendFlags(cmd)

	return run(cmd)
}
