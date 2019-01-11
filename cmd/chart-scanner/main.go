package main

import (
	"fmt"
	"log"
	"os"

	"github.com/helm/chartmuseum/pkg/config"
	"github.com/urfave/cli"
)

var (
	// Version is the semantic version (added at compile time)
	Version string

	// Revision is the git commit id (added at compile time)
	Revision string
)

func main() {
	app := cli.NewApp()
	app.Name = "chart-scanner"
	app.Version = fmt.Sprintf("%s (build %s)", Version, Revision)
	app.Usage = "checks a storage directory for evil charts"
	app.Action = cliHandler
	app.Flags = config.CLIFlags
	app.Run(os.Args)
}

func cliHandler(c *cli.Context) {
	conf := config.NewConfig()
	err := conf.UpdateFromCLIContext(c)
	if err != nil {
		log.Fatal(err)
	}
	backend := backendFromConfig(conf)
	fmt.Println(backend)
}
