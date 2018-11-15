package config

import (
	"encoding/base64"
	"io/ioutil"
	"fmt"
)

type (

	Credentials struct {
		CA string `json:"certificate-authority"`
		ClientCert string `json:"client-certificate"`
		ClientKey string `json:"client-key"`
	}

	Config struct {
		Kubeconfig     string     `json:"kubeconfig"`
		Credentials Credentials `json:"credentials"`
	}
)

const (
	ROOT_PATH        = "/root/.kube/"
	KUBECONFIG       = "kubeconfig"
	CONFIG           = "config"
	CREDENTIALS_PATH = ROOT_PATH + "credentials/"
	CA_CERT          = "ca.pem"
	CLIENT_CERT      = "client.pem"
	CLIENT_KEY       = "client-key.pem"
)

// Setup writes the kubectl config and credentials to file
func (cfg *Config) Init() error {
	// Decode and output CA cert
	ca, err := base64.StdEncoding.DecodeString(cfg.Credentials.CA)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(CREDENTIALS_PATH+CA_CERT, ca, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Wrote CA cert")

	// Decode and output client cert
	crt, err := base64.StdEncoding.DecodeString(cfg.Credentials.ClientCert)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(CREDENTIALS_PATH+CLIENT_CERT, crt, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Wrote client cert")

	// Decode and output client key
	key, err := base64.StdEncoding.DecodeString(cfg.Credentials.ClientKey)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(CREDENTIALS_PATH+CLIENT_KEY, key, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Wrote client key")

	// Decode and output kubeconfig
	config, err := base64.StdEncoding.DecodeString(cfg.Kubeconfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ROOT_PATH + CONFIG, config, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Wrote kubeconfig")

	return nil
}
