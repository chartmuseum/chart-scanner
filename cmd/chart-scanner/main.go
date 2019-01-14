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
	version string

	// Revision is the git commit id (added at compile time)
	revision string

	exitCode int
)

func main() {
	log.SetOutput(os.Stdout)
	app := cli.NewApp()
	app.Name = "chart-scanner"
	app.Version = fmt.Sprintf("%s (build %s)", version, revision[0:6])
	app.Usage = "checks a storage directory for evil charts"
	app.Flags = buildCliFlags()
	app.Action = cliHandler
	app.Run(os.Args)
}

func buildCliFlags() []cli.Flag {
	var flags []cli.Flag
	for _, flag := range config.CLIFlags {
		name := flag.GetName()
		if name == "debug" || strings.HasPrefix(name, "storage") {
			flags = append(flags, flag)
		}
	}
	sort.Sort(cli.FlagsByName(flags))
	return flags
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

	debug := conf.GetBool("debug")
	scan(backend, "", debug)

	os.Exit(exitCode)
}
