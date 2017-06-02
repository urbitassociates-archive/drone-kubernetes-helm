package commands

import (
	"os/exec"
)

func (c *Command) Dependency() error {
	cmd := exec.Command("helm", "dependency")

	if c.SubCommand != "" {
		cmd.Args = append(cmd.Args, c.SubCommand)
	}
	c.appendFlags(cmd)
	cmd.Args = append(cmd.Args, c.Chart)

	return run(cmd)
}
