package main

import (
	"log"
	"path"
	"strings"

	"github.com/chartmuseum/storage"
)

func scan(backend storage.Backend, prefix string){
	objects, _ := backend.ListObjects(prefix)
	for _, object := range objects {
		fullPath := path.Join(prefix, object.Path)
		isChartPackage := strings.HasSuffix(fullPath, ".tgz")
		if isChartPackage {
			validateChartPackage(fullPath)
		} else {
			scan(backend, fullPath)
		}
	}
}

func validateChartPackage(filePath string) {
	log.Printf("%s is valid\n", filePath)
}
