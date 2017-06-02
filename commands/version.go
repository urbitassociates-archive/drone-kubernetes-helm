package commands

import (
	"os/exec"
)

func (c *Command) Version() error {
	cmd := exec.Command("helm", "version")
	c.appendFlags(cmd)

	return run(cmd)
}
