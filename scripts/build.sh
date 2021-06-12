#!/usr/bin/env bash

set -euo pipefail

GOOS="linux" go build -ldflags='-s -w' -o bin/main github.com/paketo-buildpacks/compress-executable/cmd/main
GOOS="linux" go build -ldflags='-s -w' -o bin/helper github.com/paketo-buildpacks/compress-executable/cmd/helper

strip bin/helper bin/main
upx -q -9 bin/helper bin/main

ln -fs main bin/build
ln -fs main bin/detect
