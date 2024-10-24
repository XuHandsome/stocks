hash:=$(shell git rev-parse --short HEAD)
tag?=$(hash)

ARCH := amd64 arm64
OS := linux windows darwin

.PHONY: stocks all clean

hydra:
	go build -o stocks

all: clean
	mkdir -p build
	bash build.sh ${tag}

clean:
	rm -rf build/*
