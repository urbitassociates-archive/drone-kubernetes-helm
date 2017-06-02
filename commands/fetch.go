package commands

import (
	"os/exec"
)

func (c *Command) Fetch() error {
	cmd := exec.Command("helm", "fetch", c.Chart)
	c.appendFlags(cmd)

	return run(cmd)
}
