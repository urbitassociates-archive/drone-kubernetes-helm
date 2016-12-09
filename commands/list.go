package commands

import (
	"os"
	"os/exec"

	// "k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/services"
)

func (c *Command) List() (*services.ListReleasesResponse, error) {
	cmd := exec.Command("helm", "list")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			case "all":
				cmd.Args = append(cmd.Args, "--all")
			case "d":
				fallthrough
			case "date":
				cmd.Args = append(cmd.Args, "--date")
			case "deleted":
				cmd.Args = append(cmd.Args, "--deleted")
			case "deployed":
				cmd.Args = append(cmd.Args, "--deployed")
			case "failed":
				cmd.Args = append(cmd.Args, "--failed")
			case "m":
				fallthrough
			case "max":
				cmd.Args = append(cmd.Args, "--max", v.(string))
			case "o":
				fallthrough
			case "offset":
				cmd.Args = append(cmd.Args, "--offset", v.(string))
			case "reverse":
				cmd.Args = append(cmd.Args, "--reverse")
			case "q":
				fallthrough
			case "short":
				cmd.Args = append(cmd.Args, "--short")

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

	if c.Filter != "" {
		cmd.Args = append(cmd.Args, c.Filter)
	} else if len(c.Args) > 0 {
		for _, arg := range c.Args {
			cmd.Args = append(cmd.Args, arg)
		}
	}
	trace(cmd)

	return nil, cmd.Run()
}
