package main

// Needed in order to scan directories for local filesystem backend

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/chartmuseum/storage"
)

type LocalFilesystemBackendWithDir struct {
	*storage.LocalFilesystemBackend
}

func (b LocalFilesystemBackendWithDir) ListObjects(prefix string) ([]storage.Object, error) {
	var objects []storage.Object
	files, err := ioutil.ReadDir(path.Join(b.RootDirectory, prefix))
	if err != nil {
		if os.IsNotExist(err) {  // OK if the directory doesnt exist yet
			err = nil
		}
		return objects, err
	}
	for _, f := range files {
		object := storage.Object{Path: f.Name(), Content: []byte{}, LastModified: f.ModTime()}
		objects = append(objects, object)
	}
	return objects, nil
}

// NewLocalFilesystemBackendWithDir creates a new instance of LocalFilesystemBackendWithDir
func NewLocalFilesystemBackendWithDir(rootDirectory string) *LocalFilesystemBackendWithDir {
	absPath, err := filepath.Abs(rootDirectory)
	if err != nil {
		panic(err)
	}
	localFsBackend := &storage.LocalFilesystemBackend{RootDirectory: absPath}
	b := &LocalFilesystemBackendWithDir{localFsBackend}
	return b
}
