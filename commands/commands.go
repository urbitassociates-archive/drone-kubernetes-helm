package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

type (
	Command struct {
		Release    string                   `json:"release" yaml:"release"`
		Chart      string                   `json:"chart" yaml:"chart"`
		Flags      []map[string]interface{} `json:"flags" yaml:"flags"`
		Filter     string                   `json:"filter" yaml:"filter"`
		SubCommand string                   `json:"subcommand" yaml:"subcommand"`
		Args       []string                 `json:"args" yaml:"args"`
		Path       string                   `json:"path" yaml:"path"`
	}
)

// Invoke executes helm CLI commands, with their respective flags and values
func (c *Command) Invoke(name string) error {
	switch name {
	case "del":
		fallthrough
	case "delete":
		return c.Delete()
	case "dep":
		fallthrough
	case "dependencies":
		fallthrough
	case "dependency":
		return c.Dependency()
	case "fetch":
		return c.Fetch()
	case "get":
		return c.Get()
	case "hist":
		fallthrough
	case "history":
		return c.History()
	case "inspect":
		return c.Inspect()
	case "ls":
		fallthrough
	case "list":
		return c.List()
	case "install":
		return c.Install()
	case "lint":
		return c.Lint()
	case "repo":
		return c.Repo()
	case "rollback":
		return c.Rollback()
	case "status":
		return c.Status()
	case "test":
		return c.Test()
	case "upgrade":
		return c.Upgrade()
	case "verify":
		return c.Verify()
	case "version":
		return c.Version()
	default:
		return errors.New(fmt.Sprintf("Command '%v' not available\n", name))
	}
}

func (c *Command) appendFlags(cmd *exec.Cmd) {
	for _, flag := range c.Flags {
		for k, v := range flag {
			if v != "" {
				cmd.Args = append(cmd.Args, fmt.Sprintf("--%v=%v", k, v))
			}
		}
	}
}

func run(cmd *exec.Cmd) error {
	logrus.Print("Invoking command: ", strings.Join(cmd.Args, " "))

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	logrus.Printf("\n%v", string(out))

	return nil
}
