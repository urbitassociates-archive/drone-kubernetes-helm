package commands

import (
	"os"
	"os/exec"
)

func (c *Command) History() (interface{}, error) {
	cmd := exec.Command("helm", "history")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			// Local
			case "max":
				cmd.Args = append(cmd.Args, "--max", v.(string))

			// Global
			case "debug":
				cmd.Args = append(cmd.Args, "--debug")
			case "home":
				cmd.Args = append(cmd.Args, "--home", v.(string))
			case "host":
				cmd.Args = append(cmd.Args, "--host", v.(string))
			case "kube-context":
				cmd.Args = append(cmd.Args, "--kube-context", v.(string))
			}
		}
	}

	if len(c.Args) > 0 {
		for _, arg := range c.Args {
			cmd.Args = append(cmd.Args, arg)
		}
	} else {
		cmd.Args = append(cmd.Args, c.Release)
	}
	trace(cmd)

	return nil, cmd.Run()
}
