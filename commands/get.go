package commands

import (
	"os"
	"os/exec"
)

func (c *Command) Get() (interface{}, error) {
	cmd := exec.Command("helm", "get")
	cmd.Stderr = os.Stderr

	o := ""
	switch c.SubCommand {
	case "hooks":
		cmd.Args = append(cmd.Args, "hooks")
	case "manifest":
		cmd.Args = append(cmd.Args, "manifest")
	case "values":
		cmd.Args = append(cmd.Args, "values")
	}

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

	if len(c.Args) > 0 {
		for _, arg := range c.Args {
			cmd.Args = append(cmd.Args, arg)
		}
	} else {
		cmd.Args = append(cmd.Args, c.Release)
	}

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
