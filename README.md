# chart-scanner

[![Codefresh build status](https://g.codefresh.io/api/badges/pipeline/chartmuseum/chartmuseum%2Fchart-scanner%2Fmaster?type=cf-1)](https://g.codefresh.io/public/accounts/chartmuseum/pipelines/chartmuseum/chart-scanner/master)

## Background

This tool will attempt to detect any charts that may have been uploaded via [this vulnerability](https://helm.sh/blog/chartmuseum-security-notice-2019/index.html), affecting all versions of ChartMuseum <= 0.8.0.

## Example

The following shows detection of the test chart `evil-1.0.0.tgz` found in this repo:

```
$ git clone git@github.com:chartmuseum/chart-scanner.git
$ cd chart-scanner
$ chart-scanner --debug --storage=local --storage-local-rootdir=$(pwd)/testdata/charts
2019/01/13 17:45:17 DEBUG org1/repo1/acs-engine-autoscaler-2.2.2.tgz is valid
2019/01/13 17:45:17 DEBUG org1/repo2/aerospike-0.1.7.tgz is valid
2019/01/13 17:45:17 DEBUG org2/repo1/apm-server-0.1.0.tgz is valid
2019/01/13 17:45:17 DEBUG org2/repo2/ark-1.2.3.tgz is valid
2019/01/13 17:45:17 ERROR org2/repo2/evil-1.0.0.tgz has bad chart name "../../../../charts/org2/repo2/evil"
$ echo $?
1
```

## Installation

### CLI

Install from the latest [release artifacts](https://github.com/chartmuseum/chart-scanner/releases):
```
# Linux
curl -LO https://github.com/chartmuseum/chart-scanner/releases/download/v0.1.0/chart-scanner_0.1.0_linux_amd64.tar.gz

# macOS
curl -LO https://github.com/chartmuseum/chart-scanner/releases/download/v0.1.0/chart-scanner_0.1.0_darwin_amd64.tar.gz

# unpack, install, dispose
mkdir -p chart-scanner-install/
tar -zxf chart-scanner_0.1.0_*.tar.gz -C chart-scanner-install/
mv chart-scanner-install/chart-scanner /usr/local/bin/
rm -rf chart-scanner_0.1.0_*.tar.gz chart-scanner-install/
```

or via go get:

```
go get -u github.com/chartmuseum/chart-scanner/cmd/chart-scanner
```

Then, to run:

```
chart-scanner --help
```

### Docker Image 

A public Docker image containing the CLI is available on [Docker Hub](https://hub.docker.com/r/chartmuseum/chart-scanner):

```
docker run -it --rm chartmuseum/chart-scanner:v0.1.0 --help
```

## Usage

Command-line storage options are identical to the ones used in ChartMuseum (the package is imported and re-used).

### Using with Amazon S3

Make sure your environment is properly setup to access `my-s3-bucket`

```bash
chart-scanner --debug \
  --storage="amazon" \
  --storage-amazon-bucket="my-s3-bucket" \
  --storage-amazon-prefix="" \
  --storage-amazon-region="us-east-1"
```

You need at least the following permissions inside your IAM Policy
```yaml
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AllowListObjects",
      "Effect": "Allow",
      "Action": [
        "s3:ListBucket"
      ],
      "Resource": "arn:aws:s3:::my-s3-bucket"
    },
    {
      "Sid": "AllowObjectsCRUD",
      "Effect": "Allow",
      "Action": [
        "s3:GetObject"
      ],
      "Resource": "arn:aws:s3:::my-s3-bucket/*"
    }
  ]
}
```

### Using with Google Cloud Storage
Make sure your environment is properly setup to access `my-gcs-bucket`.

One way to do so is to set the `GOOGLE_APPLICATION_CREDENTIALS` var in your environment, pointing to the JSON file containing your service account key:
```
export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/[FILE_NAME].json"
```

More info on Google Cloud authentication can be found [here](https://cloud.google.com/docs/authentication/getting-started).

```bash
chart-scanner --debug \
  --storage="google" \
  --storage-google-bucket="my-gcs-bucket" \
  --storage-google-prefix=""
```

### Using with Microsoft Azure Blob Storage

Make sure your environment is properly setup to access `mycontainer`.

To do so, you must set the following env vars:
- `AZURE_STORAGE_ACCOUNT`
- `AZURE_STORAGE_ACCESS_KEY`

```bash
chart-scanner --debug \
  --storage="microsoft" \
  --storage-microsoft-container="mycontainer" \
  --storage-microsoft-prefix=""
```

### Using with Alibaba Cloud OSS Storage

Make sure your environment is properly setup to access `my-oss-bucket`.

To do so, you must set the following env vars:
- `ALIBABA_CLOUD_ACCESS_KEY_ID`
- `ALIBABA_CLOUD_ACCESS_KEY_SECRET`

```bash
chart-scanner --debug \
  --storage="alibaba" \
  --storage-alibaba-bucket="my-oss-bucket" \
  --storage-alibaba-prefix="" \
  --storage-alibaba-endpoint="oss-cn-beijing.aliyuncs.com"
```

### Using with Openstack Object Storage

Make sure your environment is properly setup to access `mycontainer`.

To do so, you must set the following env vars (depending on your openstack version):
- `OS_AUTH_URL`
- either `OS_PROJECT_NAME` or `OS_TENANT_NAME` or `OS_PROJECT_ID` or `OS_TENANT_ID`
- either `OS_DOMAIN_NAME` or `OS_DOMAIN_ID`
- either `OS_USERNAME` or `OS_USERID`
- `OS_PASSWORD`

```bash
chart-scanner --debug \
  --storage="openstack" \
  --storage-openstack-container="mycontainer" \
  --storage-openstack-prefix="" \
  --storage-openstack-region="myregion"
```

### Using with Oracle Cloud Infrastructure Object Storage

Make sure your environment is properly setup to access `my-ocs-bucket`.

More info on Oracle Cloud Infrastructure authentication can be found [here](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm).

```bash
chart-scanner --debug \
  --storage="oracle" \
  --storage-oracle-bucket="my-ocs-bucket" \
  --storage-oracle-prefix="" \
  --storage-oracle-compartmentid="ocid1.compartment.oc1..1234"
```

### Using with Baidu Cloud BOS Storage

Make sure your environment is properly setup to access `my-bos-bucket`.

To do so, you must set the following env vars:
- `BAIDU_CLOUD_ACCESS_KEY_ID`
- `BAIDU_CLOUD_ACCESS_KEY_SECRET`

```bash
chart-scanner --debug \
  --storage="baidu" \
  --storage-baidu-bucket="my-bos-bucket" \
  --storage-baidu-prefix="" \
  --storage-baidu-endpoint="bj.bcebos.com"
```

### Using with local filesystem storage
Make sure you have read access to `./chartstorage`.

```bash
chart-scanner --debug \
  --storage="local" \
  --storage-local-rootdir="./chartstorage"
```

### Note on environment variables
All command-line options can be specified as environment variables, which are defined by the command-line option, capitalized, with all `-`'s replaced with `_`'s.

For example, the env var `STORAGE_AMAZON_BUCKET` can be used in place of `--storage-amazon-bucket`.

## Reporting a Security Issue

If you discover a security issue in Helm or ChartMuseum, please follow the instructions found [here](https://github.com/helm/helm/blob/master/CONTRIBUTING.md#reporting-a-security-issue).
