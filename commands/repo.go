package commands

import (
	"os"
	"os/exec"
	// "k8s.io/helm/pkg/helm"
)

func (c *Command) Repo() (interface{}, error) {
	cmd := exec.Command("helm", "repo")
	cmd.Stderr = os.Stderr

	switch c.SubCommand {
	case "add":
		cmd.Args = append(cmd.Args, "add")
	case "index":
		cmd.Args = append(cmd.Args, "index")
	case "list":
		cmd.Args = append(cmd.Args, "list")
	case "remove":
		cmd.Args = append(cmd.Args, "remove")
	case "update":
		cmd.Args = append(cmd.Args, "update")
	}

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			// Local: repo add
			case "no-update":
				cmd.Args = append(cmd.Args, "--no-update")
			// Local: repo index
			case "merge":
				cmd.Args = append(cmd.Args, "--merge", v.(string))
			// Local: repo index
			case "url":
				cmd.Args = append(cmd.Args, "--url", v.(string))

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
	}

	res, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	trace(cmd)

	_, err = os.Stdout.Write(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
