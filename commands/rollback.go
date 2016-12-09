package commands

import (
	"os"
	"os/exec"

	// "k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/services"
)

func (c *Command) Rollback() (*services.RollbackReleaseResponse, error) {
	cmd := exec.Command("helm", "rollback", c.Release)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			case "dry-run":
				cmd.Args = append(cmd.Args, "--dry-run")
			case "no-hooks":
				cmd.Args = append(cmd.Args, "--no-hooks")

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

	trace(cmd)

	return nil, cmd.Run()
}
