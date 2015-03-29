WD=$(shell pwd)

default: build
	
all: clean build

build:
	mkdir -p go/{src,pkg,bin}
	GOPATH=$(WD)/go go get -d
	GOPATH="$(WD)/go" CC=clang go build -o utar -ldflags "-linkmode external -extldflags -static"
	strip utar
clean:
	rm utar
