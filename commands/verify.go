package commands

import (
	"os/exec"
)

func (c *Command) Verify() error {
	cmd := exec.Command("helm", "verify", c.Path)
	c.appendFlags(cmd)

	return run(cmd)
}
