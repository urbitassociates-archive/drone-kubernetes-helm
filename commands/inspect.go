package commands

import (
	"os/exec"
)

func (c *Command) Inspect() error {
	cmd := exec.Command("helm", "inspect")

	if c.SubCommand != "" {
		cmd.Args = append(cmd.Args, c.SubCommand)
	}
	cmd.Args = append(cmd.Args, c.Chart)
	c.appendFlags(cmd)

	return run(cmd)
}
