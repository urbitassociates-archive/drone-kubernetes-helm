package commands

import (
	"os/exec"
)

func (c *Command) List() error {
	cmd := exec.Command("helm", "list")

	if c.Filter != "" {
		cmd.Args = append(cmd.Args, c.Filter)
	}
	c.appendFlags(cmd)

	return run(cmd)
}
