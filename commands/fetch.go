package commands

import (
	"os"
	"os/exec"
)

func (c *Command) Fetch() (interface{}, error) {
	cmd := exec.Command("helm", "fetch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			case "d":
				fallthrough
			case "destination":
				cmd.Args = append(cmd.Args, "--destination", v.(string))
			case "keyring":
				cmd.Args = append(cmd.Args, "--keyring", v.(string))
			case "untar":
				cmd.Args = append(cmd.Args, "--untar")
			case "untardir":
				cmd.Args = append(cmd.Args, "--untardir", v.(string))
			case "verify":
				cmd.Args = append(cmd.Args, "--verify")
			case "version":
				cmd.Args = append(cmd.Args, "--version", v.(string))
			case "reverse":
				cmd.Args = append(cmd.Args, "--reverse")

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
		cmd.Args = append(cmd.Args, c.Chart)
	}
	trace(cmd)

	return nil, cmd.Run()
}
