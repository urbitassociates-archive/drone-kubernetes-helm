.PHONY: all clean deps fmt vet test docker

EXECUTABLE ?= drone-kubernetes-helm
IMAGE ?= mandrean/$(EXECUTABLE)
COMMIT ?= $(shell git rev-parse --short HEAD)
TAG ?= $(git describe --exact-match --tags HEAD)

LDFLAGS = -X "main.buildCommit=$(COMMIT)"
PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: deps build test

clean:
	go clean -i ./...

deps:
	glide install

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

docker:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)'
	docker build --build-arg K8S_VERSION=1.4.12 --build-arg HELM_VERSION=2.4.2 --rm --no-cache -t $(IMAGE):v0.3.0-k8s-v1.4.12-helm-v2.4.2 .
	docker build --build-arg K8S_VERSION=1.5.7  --build-arg HELM_VERSION=2.4.2 --rm --no-cache -t $(IMAGE):v0.3.0-k8s-v1.5.7-helm-v2.4.0 .
	docker build --build-arg K8S_VERSION=1.6.4  --build-arg HELM_VERSION=2.4.2 --rm --no-cache -t $(IMAGE):v0.3.0-k8s-v1.6-4-helm-v2.4.2 .

$(EXECUTABLE): $(wildcard *.go)
	go build -ldflags '-s -w $(LDFLAGS)'

build: $(EXECUTABLE)
