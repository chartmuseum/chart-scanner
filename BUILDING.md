# Developer Guide

## Building binary

Use `make test` to build all platofrm binaries to the `bin/` directory.

Mac:

```
# builds to bin/darwin/amd64/chart-scanner
make build-mac
```

Linux:

```
# builds to bin/linux/amd64/chart-scanner
make build-linux
```

Windows:

```
# builds to bin/windows/amd64/chart-scanner.exe
make build-windows
```

## Cleaning workspace

To remove all files not manged by git, run `make clean` (be careful!)

## Adding/updating dependencies

Requires [dep](https://golang.github.io/dep/).

Add new dependencies directly to `Gopkg.toml` and run `dep ensure`.

To update all dependencies, run `make update-deps`.

Afterwords, run `make fix-deps`, which will replace one vendored source file.

Please check in all changes in the `vendor/` directory.

## Cutting a new release

Example of releasing `v0.1.0`:
```
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

A Codefresh pipeline will pick up the GitHub tag event
and run [.codefresh/release.yml](.codefresh/release.yml).

This will result in running [goreleaser](https://goreleaser.com/)
to upload release artiacts, as well as push a tag to Docker Hub for
the image `jdolitsky/chart-scanner`.
