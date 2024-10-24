#!/bin/bash

TAG=$1

for arch in amd64 arm64; do
    for os in linux windows darwin; do
        build_dir=stocks-"${os}"-"${arch}"-"${TAG}"
        mkdir -p "${build_dir}"
        GOOS="${os}" GOARCH="${arch}" CGO_ENABLED=0 go build -o "${build_dir}"/
        cp -Raf config.yaml "${build_dir}"/
        tar zcvf build/"${build_dir}".tgz "${build_dir}"
        rm -rf "${build_dir}"
    done
done