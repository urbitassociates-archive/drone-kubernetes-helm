package commands

import (
	"os/exec"
)

func (c *Command) Rollback() error {
	cmd := exec.Command("helm", "rollback", c.Release)
	c.appendFlags(cmd)

	if len(c.Args) > 0 {
		for _, arg := range c.Args {
			cmd.Args = append(cmd.Args, arg)
		}
	}

	return run(cmd)
}
