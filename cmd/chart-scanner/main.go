package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

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
	log.SetOutput(os.Stdout)
	app := cli.NewApp()
	app.Name = "chart-scanner"
	app.Version = fmt.Sprintf("%s (build %s)", Version, Revision)
	app.Usage = "checks a storage directory for evil charts"
	app.Action = cliHandler
	var flags []cli.Flag
	for _, flag := range config.CLIFlags {
		if strings.HasPrefix(flag.GetName(), "storage") {
			flags = append(flags, flag)
		}
	}
	app.Flags = flags
	sort.Sort(cli.FlagsByName(app.Flags))
	app.Run(os.Args)
}

func cliHandler(c *cli.Context) {
	conf := config.NewConfig()
	err := conf.UpdateFromCLIContext(c)
	if err != nil {
		log.Fatal(err)
	}

	backend := backendFromConfig(conf)

	// First make sure we have access to the root storage dir
	err = check(backend)
	if err != nil {
		log.Fatal(err)
	}

	scan(backend, "")
}
