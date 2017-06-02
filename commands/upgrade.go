package commands

import (
	"os/exec"
)

func (c *Command) Upgrade() error {
	cmd := exec.Command("helm", "upgrade", c.Release, c.Chart)
	c.appendFlags(cmd)

	return run(cmd)
}
