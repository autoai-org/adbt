#!/usr/bin/env bash
package_name="adbt"
platforms=("windows/amd64" "darwin/amd64" "linux/amd64" "linux/arm64" "linux/mips" "freebsd/amd64")
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    echo 'Building ADBT for '$GOOS'/'$GOARCH'...'
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}"  -o 'dist/'$output_name
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done