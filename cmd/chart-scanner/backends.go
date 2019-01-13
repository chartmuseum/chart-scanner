package main

// Pretty much copy-paste from chartmuseum cmd/

import (
	"fmt"
	"log"
	"strings"

	"github.com/chartmuseum/storage"
	"github.com/helm/chartmuseum/pkg/config"
)

func backendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.backend"})

	var backend storage.Backend

	storageFlag := strings.ToLower(conf.GetString("storage.backend"))
	switch storageFlag {
	case "local":
		backend = localBackendFromConfig(conf)
	case "amazon":
		backend = amazonBackendFromConfig(conf)
	case "google":
		backend = googleBackendFromConfig(conf)
	case "oracle":
		backend = oracleBackendFromConfig(conf)
	case "microsoft":
		backend = microsoftBackendFromConfig(conf)
	case "alibaba":
		backend = alibabaBackendFromConfig(conf)
	case "openstack":
		backend = openstackBackendFromConfig(conf)
	default:
		log.Fatal("Unsupported storage backend: ", storageFlag)
	}

	return backend
}

func localBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.local.rootdir"})
	return storage.Backend(NewLocalFilesystemBackendWithDir(
		conf.GetString("storage.local.rootdir"),
	))
}

func amazonBackendFromConfig(conf *config.Config) storage.Backend {
	// If using alternative s3 endpoint (e.g. Minio) default region to us-east-1
	if conf.GetString("storage.amazon.endpoint") != "" && conf.GetString("storage.amazon.region") == "" {
		conf.Set("storage.amazon.region", "us-east-1")
	}
	crashIfConfigMissingVars(conf, []string{"storage.amazon.bucket", "storage.amazon.region"})
	return storage.Backend(storage.NewAmazonS3Backend(
		conf.GetString("storage.amazon.bucket"),
		conf.GetString("storage.amazon.prefix"),
		conf.GetString("storage.amazon.region"),
		conf.GetString("storage.amazon.endpoint"),
		conf.GetString("storage.amazon.sse"),
	))
}

func googleBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.google.bucket"})
	return storage.Backend(storage.NewGoogleCSBackend(
		conf.GetString("storage.google.bucket"),
		conf.GetString("storage.google.prefix"),
	))
}

func oracleBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.oracle.bucket", "storage.oracle.compartmentid"})
	return storage.Backend(storage.NewOracleCSBackend(
		conf.GetString("storage.oracle.bucket"),
		conf.GetString("storage.oracle.prefix"),
		conf.GetString("storage.oracle.region"),
		conf.GetString("storage.oracle.compartmentid"),
	))
}

func microsoftBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.microsoft.container"})
	return storage.Backend(storage.NewMicrosoftBlobBackend(
		conf.GetString("storage.microsoft.container"),
		conf.GetString("storage.microsoft.prefix"),
	))
}

func alibabaBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.alibaba.bucket"})
	return storage.Backend(storage.NewAlibabaCloudOSSBackend(
		conf.GetString("storage.alibaba.bucket"),
		conf.GetString("storage.alibaba.prefix"),
		conf.GetString("storage.alibaba.endpoint"),
		conf.GetString("storage.alibaba.sse"),
	))
}

func openstackBackendFromConfig(conf *config.Config) storage.Backend {
	crashIfConfigMissingVars(conf, []string{"storage.openstack.container", "storage.openstack.region"})
	return storage.Backend(storage.NewOpenstackOSBackend(
		conf.GetString("storage.openstack.container"),
		conf.GetString("storage.openstack.prefix"),
		conf.GetString("storage.openstack.region"),
		conf.GetString("storage.openstack.cacert"),
	))
}

func crashIfConfigMissingVars(conf *config.Config, vars []string) {
	missing := []string{}
	for _, v := range vars {
		if conf.GetString(v) == "" {
			flag := config.GetCLIFlagFromVarName(v)
			missing = append(missing, fmt.Sprintf("--%s", flag))
		}
	}
	if len(missing) > 0 {
		log.Fatal("Missing required flags(s): ", strings.Join(missing, ", "))
	}
}
