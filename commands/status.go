package commands

import (
	"os"
	"os/exec"

	// "k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/services"
)

func (c *Command) Status() (*services.GetReleaseStatusResponse, error) {
	cmd := exec.Command("helm", "status", c.Release)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			case "revision":
				cmd.Args = append(cmd.Args, "--revision", v.(string))

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

	trace(cmd)

	return nil, cmd.Run()
}
