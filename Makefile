.PHONY: build test clean docker

GO=CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/device-service/device-service
.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)

GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-service.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: $(MICROSERVICES)
	$(GO) install -tags=safe

build-rasp: 
	OOS=linux GOARCH=arm GOARM=5 $(GO) build $(GOFLAGS) -o $@ ./cmd/device-service 
	
cmd/device-service/device-service:
	OOS=linux GOARCH=arm GOARM=5 $(GO) build $(GOFLAGS) -o $@ ./cmd/device-service

docker:
	docker build \
		-f example/cmd/device-simple/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-device-sdk-simple:$(GIT_SHA) \
		-t edgexfoundry/docker-device-sdk-simple:$(VERSION)-dev \
		.

test:
	$(GO) vet ./...
	gofmt -l .
	$(GO) test -coverprofile=coverage.out ./...

clean:
	rm -f $(MICROSERVICES)
