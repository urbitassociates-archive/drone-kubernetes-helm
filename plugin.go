package main

import (
	"io/ioutil"

	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/mandrean/drone-kubernetes-helm/commands"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"strings"
)

type (
	Cluster struct {
		Name    string      `json:"name" yaml:"name"`
		Cluster interface{} `json:"cluster" yaml:"cluster"`
	}

	User struct {
		Name string      `json:"name" yaml:"name"`
		User interface{} `json:"user" yaml:"user"`
	}

	Context struct {
		Name    string `json:"name" yaml:"name"`
		Context struct {
			Cluster string      `json:"cluster" yaml:"cluster"`
			User    interface{} `json:"user" yaml:"user"`
		} `json:"context" yaml:"context"`
	}

	KubeConfig struct {
		ApiVersion     string      `json:"apiVersion" yaml:"apiVersion"`
		Kind           string      `json:"kind" yaml:"kind"`
		Clusters       []Cluster   `json:"clusters" yaml:"clusters"`
		Users          []User      `json:"users" yaml:"users"`
		Contexts       []Context   `json:"contexts" yaml:"contexts"`
		CurrentContext string      `json:"current-context" yaml:"current-context"`
		Preferences    interface{} `json:"preferences" yaml:"preferences"`
	}

	// Plugin defines the Docker plugin parameters.
	Plugin struct {
		Context      *cli.Context
		KubeConfig   KubeConfig                    `json:"kube-config" yaml:"kube-config"`
		HelmCommands []map[string]commands.Command `json:"helm-commands" yaml:"helm-commands"`
	}
)

// Exec executes the plugin step
func (p Plugin) Exec() error {
	if err := p.writeKubeConfig(); err != nil {
		return err
	}

	// loop through commands
	for _, command := range p.HelmCommands {
		for k, v := range command {
			if err := v.Invoke(k); err != nil {
				return err
			}
		}
	}

	return nil
}

func NewKubeConfig() KubeConfig {
	return KubeConfig{
		ApiVersion: "v1",
		Kind:       "Config",
	}
}

// createKubeconfig creates the kubeconfig file
func (p Plugin) writeKubeConfig() error {
	var homePath = strings.ToLower(os.Getenv("HOME"))
	var kubeConfigPath = path.Join(homePath, "/.kube/config")

	logrus.Debug("Marshaling kube config...")
	y, err := yaml.Marshal(p.KubeConfig)
	if err != nil {
		logrus.Fatalln(err)
		return errors.New("Could not marshal kubeconfig YAML.")
	}
	logrus.Debug("Marshaled kube config.")

	logrus.Debug("Writing kube config...")
	err = ioutil.WriteFile(kubeConfigPath, y, 0644)
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Debugf("Wrote kube config:\n%v", string(y))

	return nil
}
