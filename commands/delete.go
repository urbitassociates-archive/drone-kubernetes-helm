package commands

import (
	"os/exec"
)

func (c *Command) Delete() error {
	cmd := exec.Command("helm", "delete", c.Release)
	c.appendFlags(cmd)

	return run(cmd)
}
