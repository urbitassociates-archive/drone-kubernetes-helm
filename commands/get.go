package commands

import (
	"os"
	"os/exec"
)

func (c *Command) Get() error {
	cmd := exec.Command("helm", "get")
	cmd.Stderr = os.Stderr

	if c.SubCommand != "" {
		cmd.Args = append(cmd.Args, c.SubCommand)
	}

	c.appendFlags(cmd)

	if len(c.Args) > 0 {
		for _, arg := range c.Args {
			cmd.Args = append(cmd.Args, arg)
		}
	} else {
		cmd.Args = append(cmd.Args, c.Release)
	}

	return run(cmd)
}
