package main

import (
	"log"

	"github.com/chartmuseum/storage"
)

func scanBackend(backend storage.Backend) error {
	objects, err := backend.ListObjects("")
	if err != nil {
		return err
	}
	log.Println(len(objects))
	return nil
}
