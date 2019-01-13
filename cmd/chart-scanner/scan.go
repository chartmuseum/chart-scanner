package main

import (
	"log"
	"path"
	"strings"

	"github.com/chartmuseum/storage"
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
			validateChartPackage(fullPath, debug)
		} else {
			scan(backend, fullPath, debug)
		}
	}
}

func validateChartPackage(filePath string, debug bool) {
	if debug {
		log.Printf("DEBUG %s is valid\n", filePath)
	}
}
