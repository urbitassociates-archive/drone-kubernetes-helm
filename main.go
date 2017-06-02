package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	legacyPlugin "github.com/drone/drone-go/plugin" // https://github.com/drone/drone-go/tree/eaa41f7836a191224ec5702d2458db07882ec269
	"github.com/imdario/mergo"
	"github.com/joho/godotenv"
	"github.com/mandrean/drone-kubernetes-helm/commands"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var (
	build = "0" // build number set at compile-time
)

func main() {
	// Load env-file if it exists first
	if env := os.Getenv("PLUGIN_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	// determine if in debug mode
	if (os.Getenv("DEBUG") == "true") || (os.Getenv("PLUGIN_DEBUG") == "true") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	app := cli.NewApp()
	app.Name = "kubernetes-helm plugin"
	app.Usage = "kubernetes-helm plugin"
	app.Action = run
	app.Version = fmt.Sprintf("0.3.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "kube-config",
			Usage:  "Kubernetes Kubectl/Helm configuration",
			EnvVar: "PLUGIN_KUBE-CONFIG, PLUGIN_KUBE_CONFIG, KUBE-CONFIG, KUBE_CONFIG",
		},
		cli.StringFlag{
			Name:   "helm-commands",
			Usage:  "helm commands to execute in order",
			EnvVar: "PLUGIN_HELM-COMMANDS, PLUGIN_HELM_COMMANDS, HELM-COMMANDS, HELM_COMMANDS",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "run the plugin in debug mode",
			EnvVar: "PLUGIN_DEBUG, DEBUG",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{}
	plugin.Context = c

	if !plugin.isLegacy() {
		plugin.KubeConfig = unmarshalKubeConfig(c.String("kube-config"))
		plugin.HelmCommands = unmarshalCommands(c.String("helm-commands"))
	}

	return plugin.Exec()
}

func hasStdInput() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func (p *Plugin) isLegacy() bool {
	if !hasStdInput() {
		return false
	}

	logrus.Debug("Parsing legacy vargs...")
	legacyPlugin.Param("vargs", p)
	if err := legacyPlugin.Parse(); err != nil {
		logrus.Debugf("Error parsing legacy vargs: %#v. Trying normal mode.", err)
		return false
	}

	if p == nil || p.HelmCommands == nil {
		logrus.Debug("Legacy vargs empty. Trying normal mode.")
		return false
	}

	logrus.Debug("Running in legacy mode.")

	return true
}

func unmarshalKubeConfig(jsonKubeConfig string) KubeConfig {
	var cfg KubeConfig

	logrus.Debug("Unmarshaling kube config...")

	if err := yaml.Unmarshal([]byte(jsonKubeConfig), &cfg); err != nil {
		logrus.Fatal(err)
	}

	if err := mergo.Merge(&cfg, NewKubeConfig()); err != nil {
		logrus.Fatal(err)
	}

	logrus.Debug("Unmarshaled kube config.")

	return cfg
}

func unmarshalCommands(jsonCommands string) []map[string]commands.Command {
	var cmds []map[string]commands.Command

	logrus.Debug("Unmarshaling helm commands...")

	if err := yaml.Unmarshal([]byte(jsonCommands), &cmds); err != nil {
		logrus.Fatal(err)
	}

	logrus.Debug("Unmarshaled helm commands.")

	return cmds
}
