package commands

import (
	"os/exec"
)

func (c *Command) Install() error {
	cmd := exec.Command("helm", "install", c.Chart)
	c.appendFlags(cmd)

	return run(cmd)
}
