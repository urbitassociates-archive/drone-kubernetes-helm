package main

import (
	"log"
	"fmt"
	"os"
	"os/exec"
	"strings"

	// "k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/services"
	// "github.com/drone/drone-go/plugin" // https://github.com/drone/drone-go/tree/eaa41f7836a191224ec5702d2458db07882ec269
	"github.com/drone/drone-plugin-go/plugin"
	"strconv"
	"encoding/base64"
	"io/ioutil"
)

type HelmClient interface {
	list() (*services.ListReleasesResponse, error)
	install() (*services.InstallReleaseResponse, error)
	delete() (*services.UninstallReleaseResponse, error)
	upgrade() (*services.UpdateReleaseResponse, error)
	rollback() (*services.RollbackReleaseResponse, error)
	status() (*services.GetReleaseStatusResponse, error)
	content() (*services.GetReleaseContentResponse, error)
	history() (*services.GetHistoryResponse, error)
	version() (*services.GetVersionResponse, error)
}

type (
	CommandName string

	Credentials struct {
		CA   string `json:"certificate-authority"` // cat ca.pem | base64
		Cert string `json:"client-certificate"` // cat client.pem | base64
		Key  string `json:"client-key"` // cat client-key.pem | base64
	}

	Config struct {
		KubeConfig  string `json:"kubeconfig"` // $ cat ~/.kube/config | base64
		Credentials Credentials `json:"credentials"`
	}

	Helm struct {
		Config   Config `json:"config"`
		Commands []map[string]Command `json:"commands"`
	}

	Command struct {
		Release    string `json:"release"` // "dev-store-api"
		Chart      string `json:"chart"`// "store-api"
		Flags      []map[string]interface{} `json:"flags"` // ["namespace"] = "kube-system"
		Filter     string `json:"filter"` // "dev"
		SubCommand string `json:"command"` // "values"
	}
)

var (
	buildCommit string
)

const (
	KUBE_PATH = "/root/.kube/"
	KUBE_CONFIG = "config"
	KUBE_CREDENTIALS_PATH = KUBE_PATH + "credentials/"
	KUBE_CA_CERT = "ca.pem"
	KUBE_CLIENT_CERT = "client.pem"
	KUBE_CLIENT_KEY = "client-key.pem"
)

func main() {
	fmt.Printf("Drone Kubernetes Helm Plugin built from %s\n", buildCommit)

	workspace := plugin.Workspace{}
	build := plugin.Build{}
	vargs := Helm{}

	plugin.Param("workspace", &workspace)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	err := vargs.Config.setup()
	if err != nil {
		log.Fatalln(err)
	}

	// loop through commands
	for _, command := range vargs.Commands {
		for k, v := range command {
			fmt.Println("Command:", k)
			r, e := v.Invoke(k)
			if e != nil {
				fmt.Fprintln(os.Stderr)
				fmt.Println(r)
				log.Fatalln(e)
			}

			fmt.Fprintln(os.Stdout)
		}
	}
}

// Invoke executes helm CLI commands, with their respective flags and values
func (c *Command) Invoke(name string) (interface{}, error) {
	switch name {
	case "fetch":
		return c.Fetch()
	case "get":
		return c.Get()
	case "history":
		return c.History()
	case "list":
		return c.List()
	case "install":
		return c.Install()
	case "delete":
		return c.Delete()
	case "upgrade":
		return c.Upgrade()
	case "rollback":
		return c.Rollback()
	case "status":
		return c.Status()
	case "version":
		return c.Version()
	default:
		return c.Version()
	}
}

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

	cmd.Args = append(cmd.Args, c.Chart)
	trace(cmd)

	return nil, cmd.Run()
}

func (c *Command) Get() (interface{}, error) {
	cmd := exec.Command("helm", "get")
	cmd.Stderr = os.Stderr

	o := ""
	cmd.Args = append(cmd.Args, c.SubCommand)

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			// Local
			case "all":
				cmd.Args = append(cmd.Args, "--all")
			case "revision":
				cmd.Args = append(cmd.Args, "--revision", v.(string))
			case "output":
				o = v.(string)

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

	cmd.Args = append(cmd.Args, c.Release)
	res, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if o != "" {
		f, err := os.Create(o)
		if err != nil {
			return nil, err
		}

		_, err = f.Write(res)
		if err != nil {
			return nil, err
		}

		err = f.Close()
		if err != nil {
			return nil, err
		}
	}
	trace(cmd)

	_, err = os.Stdout.Write(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

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

	cmd.Args = append(cmd.Args, c.Release)
	trace(cmd)

	return nil, cmd.Run()
}

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
	}

	trace(cmd)

	return nil, cmd.Run()
}

func (c *Command) Install() (*services.InstallReleaseResponse, error) {
	ns := "default"
	cmd := exec.Command("helm", "install", c.Chart)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			case "dry-run":
				cmd.Args = append(cmd.Args, "--dry-run")
			case "keyring":
				cmd.Args = append(cmd.Args, "--keyring", v.(string))
			case "n":
				fallthrough
			case "name":
				rn := ""
				if c.Release != "" {
					rn = c.Release
				}
				if v != "" {
					rn = v.(string)
				}
				cmd.Args = append(cmd.Args, "--name", rn)
			case "name-template":
				cmd.Args = append(cmd.Args, "--name-template", v.(string))
			case "namespace":
				ns = v.(string)
			case "no-hooks":
				cmd.Args = append(cmd.Args, "--no-hooks")
			case "replace":
				cmd.Args = append(cmd.Args, "--replace")
			case "set":
				cmd.Args = append(cmd.Args, "--set", v.(string))
			case "f":
				fallthrough
			case "values":
				cmd.Args = append(cmd.Args, "--values", v.(string))
			case "verify":
				cmd.Args = append(cmd.Args, "--verify")
			case "version":
				cmd.Args = append(cmd.Args, "--version", v.(string))

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

	cmd.Args = append(cmd.Args, "--namespace", ns)
	trace(cmd)

	rev1, err := latestRevision(c.Release)
	if err != nil {
		return nil, err
	}

	err = cmd.Run()
	if err != nil {
		rollbackRevChange(c.Release, rev1)
	}

	return nil, err
}

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

func (c *Command) Upgrade() (*services.UpdateReleaseResponse, error) {
	ns := "default"
	cmd := exec.Command("helm", "upgrade", c.Release, c.Chart)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			// Local
			case "disable-hooks":
				cmd.Args = append(cmd.Args, "--disable-hooks")
			case "dry-run":
				cmd.Args = append(cmd.Args, "--dry-run")
			case "i":
				fallthrough
			case "install":
				cmd.Args = append(cmd.Args, "--install")
			case "keyring":
				cmd.Args = append(cmd.Args, "--keyring", v.(string))
			case "namespace":
				ns = v.(string)
			case "set":
				cmd.Args = append(cmd.Args, "--set", v.(string))
			case "f":
				fallthrough
			case "values":
				cmd.Args = append(cmd.Args, "--values", v.(string))
			case "verify":
				cmd.Args = append(cmd.Args, "--verify")
			case "version":
				cmd.Args = append(cmd.Args, "--version", v.(string))

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

	cmd.Args = append(cmd.Args, "--namespace", ns)
	trace(cmd)

	rev1, err := latestRevision(c.Release)
	if err != nil {
		return nil, err
	}

	err = cmd.Run()
	if err != nil {
		rollbackRevChange(c.Release, rev1)
	}

	return nil, err
}

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

	trace(cmd)

	return nil, cmd.Run()
}

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

func (c *Command) Version() (*services.GetVersionResponse, error) {
	cmd := exec.Command("helm", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			case "c":
				fallthrough
			case "client":
				cmd.Args = append(cmd.Args, "--client")
			case "s":
				fallthrough
			case "server":
				cmd.Args = append(cmd.Args, "--server")

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

// Trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}

// LatestRevision returns the current latest
func latestRevision(release string) (int, error) {
	res, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("helm history %v | tail -n1 | cut -d \" \" -f 1", release),
	).Output()
	if err != nil {
		return -1, err
	}
	s := string(res[:])
	if s == "" {
		return 0, nil
	}
	revision, err := strconv.Atoi(strings.Trim(s, "\n"))
	if err != nil {
		return -1, err
	}

	return revision, nil
}

// Rollback rolls back the release to revision rev1 if the current revision differs from rev1
func rollbackRevChange(release string, rev1 int) error {
	rev2, err := latestRevision(release)
	if err != nil {
		return err
	}

	if rev2 != rev1 {
		res, err := exec.Command("helm", "rollback", release, strconv.Itoa(rev1)).Output()
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(res)
		if err != nil {
			return err
		}
	}

	return nil
}

// Setup writes the kubectl config and credentials to file
func (cfg *Config) setup() error {
	kubeCfg, err := base64.StdEncoding.DecodeString(cfg.KubeConfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(KUBE_PATH + KUBE_CONFIG, kubeCfg, 0644)
	if err != nil {
		return err
	}

	ca, err := base64.StdEncoding.DecodeString(cfg.Credentials.CA)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(KUBE_CREDENTIALS_PATH + KUBE_CA_CERT, ca, 0644)
	if err != nil {
		return err
	}

	crt, err := base64.StdEncoding.DecodeString(cfg.Credentials.Cert)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(KUBE_CREDENTIALS_PATH + KUBE_CLIENT_CERT, crt, 0644)
	if err != nil {
		return err
	}

	key, err := base64.StdEncoding.DecodeString(cfg.Credentials.Key)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(KUBE_CREDENTIALS_PATH + KUBE_CLIENT_KEY, key, 0644)
	if err != nil {
		return err
	}

	return nil
}
