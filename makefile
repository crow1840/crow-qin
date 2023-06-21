.PHONY: dockerbuild pushimage

GOPATH:=$(shell go env GOPATH)
ROOT_DIR = $(CURDIR)
SRC_DIR = $(ROOT_DIR)
BIN_DIR = $(ROOT_DIR)/bin
BIN_DIR_LOCAL = $(BIN_DIR)/local
BIN_DIR_LINUX = $(BIN_DIR)/linux_amd64
DOCKER_DIR = $(ROOT_DIR)/docker
BIN_DIR_DOCKER = $(DOCKER_DIR)/bin
TOOLS_DIR = $(ROOT_DIR)/tools

CURR_TIME = $(shell date "+%Y-%m-%d %H:%M:%S")
HOST_NAME = $(shell hostname)
GO_VERSION = $(shell go version)
GIT_LOG = $(shell git log -1 | sed s/\'/\"/g)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
var:=
DOCKER_HUB:=
DOCKER_HUB_NS:=crow
VERSION_SUFFIX=
webroot:=
DOCKER_IMAGE_TAG=latest
DOCKER_IMAGE_SHORT_VER=-v3

VERSION=$(shell cat VERSION)

LDFLAGS = "-X 'router.BuildTime=${CURR_TIME}' \
-X 'crow-qin/router.BuildHost=${HOST_NAME}' \
-X 'crow-qin/router.GOVersion=${GO_VERSION}' \
-X 'crow-qin/router.GitLog=${GIT_LOG}' \
-X 'crow-qin/router.GitBranch=${GIT_BRANCH}' \
-X 'crow-qin/router.CrowVersion=${VERSION}'"

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -installsuffix cgo -ldflags $(LDFLAGS) -o ./crow-qin

flags:
	echo  $(LDFLAGS)

pushimage:
	docker build -t crow/crow-qin .
#	docker push crow/crow-qin

dockerbuild: build pushimage