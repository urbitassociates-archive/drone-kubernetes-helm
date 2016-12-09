package commands

import (
	"os"
	"os/exec"

	// "k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/services"
)

func (c *Command) Delete() (*services.UninstallReleaseResponse, error) {
	cmd := exec.Command("helm", "delete", c.Release)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			// Local
			case "dry-run":
				cmd.Args = append(cmd.Args, "--dry-run")
			case "no-hooks":
				cmd.Args = append(cmd.Args, "--no-hooks")
			case "purge":
				cmd.Args = append(cmd.Args, "--purge")

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
