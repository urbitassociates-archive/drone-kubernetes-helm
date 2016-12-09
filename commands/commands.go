package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/pkg/errors"
)

type (
	Command struct {
		Release    string                   `json:"release"`    // "dev-store-api"
		Chart      string                   `json:"chart"`      // "store-api"
		Flags      []map[string]interface{} `json:"flags"`      // ["namespace"] = "kube-system"
		Filter     string                   `json:"filter"`     // "dev"
		SubCommand string                   `json:"subcommand"` // "values"
		Args       []string                 `json:"args"`       // ["name", "url"]
	}
)

// Invoke executes helm CLI commands, with their respective flags and values
func (c *Command) Invoke(name string) (interface{}, error) {
	switch name {
	case "fetch":
		return c.Fetch()
	case "get":
		return c.Get()
	case "hist":
		fallthrough
	case "history":
		return c.History()
	case "ls":
		fallthrough
	case "list":
		return c.List()
	case "install":
		return c.Install()
	case "del":
		fallthrough
	case "delete":
		return c.Delete()
	case "upgrade":
		return c.Upgrade()
	case "repo":
		return c.Repo()
	case "rollback":
		return c.Rollback()
	case "status":
		return c.Status()
	case "version":
		return c.Version()
	default:
		return nil, errors.New(fmt.Sprintf("Command '%v' not available\n", name))
	}
}

// Trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
