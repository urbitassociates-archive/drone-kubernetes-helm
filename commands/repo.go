package commands

import (
	"os/exec"
)

func (c *Command) Repo() error {
	cmd := exec.Command("helm", "repo")

	if c.SubCommand != "" {
		cmd.Args = append(cmd.Args, c.SubCommand)
	}
	c.appendFlags(cmd)

	return run(cmd)
}
