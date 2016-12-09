package config

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"text/template"
	"fmt"
)

type (
	Kubeconfig struct {
		APIServer string `json:"api_server"`
	}

	// X509 Client Certificate authentication
	X509 struct {
		CA   string `json:"certificate_authority"` // cat ca.pem | base64
		Cert string `json:"client_certificate"`    // cat client.pem | base64
		Key  string `json:"client_key"`            // cat client-key.pem | base64
	}

	Authentication struct {
		ClientCert X509 `json:"client_cert"`
		X509       X509 `json:"x509"`
	}

	Config struct {
		Kubeconfig     Kubeconfig     `json:"kubeconfig"`
		Authentication Authentication `json:"authentication"`
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
	ca, err := base64.StdEncoding.DecodeString(cfg.Authentication.X509.CA)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(CREDENTIALS_PATH+CA_CERT, ca, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Wrote CA cert")

	// Decode and output client cert
	crt, err := base64.StdEncoding.DecodeString(cfg.Authentication.X509.Cert)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(CREDENTIALS_PATH+CLIENT_CERT, crt, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Wrote client cert")

	// Decode and output client key
	key, err := base64.StdEncoding.DecodeString(cfg.Authentication.X509.Key)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(CREDENTIALS_PATH+CLIENT_KEY, key, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Wrote client key")

	// Output kubeconfig
	t, err := template.ParseFiles(ROOT_PATH + KUBECONFIG)
	if err != nil {
		return err
	}
	f, err := os.Create(ROOT_PATH + CONFIG)
	if err != nil {
		return err
	}
	err = t.Execute(f, cfg.Kubeconfig)
	if err != nil {
		return err
	}
	f.Close()
	fmt.Println("Wrote kubeconfig")

	return nil
}
