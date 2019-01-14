package main

import (
	"bytes"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/chartmuseum/storage"
	"k8s.io/helm/pkg/chartutil"
)

func check(backend storage.Backend) error {
	_, err := backend.ListObjects("")
	return err
}

func scan(backend storage.Backend, prefix string, debug bool) {
	objects, _ := backend.ListObjects(prefix)
	for _, object := range objects {
		fullPath := path.Join(prefix, object.Path)
		isChartPackage := strings.HasSuffix(fullPath, ".tgz")
		if isChartPackage {
			validateChartPackage(backend, fullPath, debug)
		} else {
			scan(backend, fullPath, debug)
		}
	}
}

func validateChartPackage(backend storage.Backend, filePath string, debug bool) {
	object, err := backend.GetObject(filePath)
	if err != nil {
		log.Printf("ERROR %s could not be retrieved\n", filePath)
		exitCode = 1
		return
	}
	chart, err := chartutil.LoadArchive(bytes.NewBuffer(object.Content))
	if err != nil {
		log.Printf("ERROR %s could not be loaded as a chart\n", filePath)
		exitCode = 1
		return
	}

	fileBaseName := filepath.Base(filePath)
	name := chart.Metadata.Name
	version := chart.Metadata.Version

	// the actual validation occurs here
	if strings.ContainsAny(name, "/\\") ||
		strings.ContainsAny(version, "/\\") ||
		fileBaseName != fmt.Sprintf("%s-%s.tgz", name, version) {

		log.Printf("ERROR %s has bad chart name \"%s\"\n", filePath, name)
		exitCode = 1
		return
	}

	if debug {
		log.Printf("DEBUG %s is valid\n", filePath)
	}
}
