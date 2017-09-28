#
# dbs-samples Makefile
#
DIST = dist

VERSION = 0.1.0
SYSARCH = amd64
STAGING = devel
GITHASH = $(shell git rev-parse --short HEAD)
VSUFFIX = banjo-$(VERSION)-$(SYSARCH)-$(STAGING)

PROJECT = github.com/smakogon/dbs-samples

all: build

init:
	curl https://glide.sh/get | sh

deps:
	glide up

build:
	GOOS=linux CGO_ENABLED=0 GOARCH=$(SYSARCH) \
		go build -a -o $(DIST)/bin/$(VSUFFIX) src/main.go

dist: build
	tar cvzf $(VSUFFIX).tgz dist/*

clean:
	rm -f *.tgz
	rm -f $(DIST)/bin/*

.PHONY: all init deps build dist clean
