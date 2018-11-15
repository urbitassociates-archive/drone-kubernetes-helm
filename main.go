package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/drone/drone-go/plugin" // https://github.com/drone/drone-go/tree/eaa41f7836a191224ec5702d2458db07882ec269
	"github.com/drone/drone-plugin-go/plugin"

	"github.com/urbitassociates/drone-kubernetes-helm/commands"
	"github.com/urbitassociates/drone-kubernetes-helm/config"
)

type (
	Helm struct {
		Config   config.Config                 `json:"config"`
		Commands []map[string]commands.Command `json:"commands"`
	}
)

var (
	buildCommit string
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

	err := vargs.Config.Init()
	if err != nil {
		log.Fatalln(err)
	}

	// loop through commands
	for _, command := range vargs.Commands {
		for k, v := range command {
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
