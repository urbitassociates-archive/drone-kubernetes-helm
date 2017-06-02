package commands

import (
	"os/exec"
)

func (c *Command) Lint() error {
	cmd := exec.Command("helm", "lint", c.Path)
	c.appendFlags(cmd)

	return run(cmd)
}
