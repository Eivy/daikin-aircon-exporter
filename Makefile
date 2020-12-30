VERSION ?= devel
IMAGE_TAG ?= devel
IMAGE_ARCH=
.PHONY: build/daikin-aircon-exporter
build/daikin-aircon-exporter:
	mkdir -p build
	go build -o $@ -ldflags "-X github.com/eivy/daikin-aircon-exporter/exporter.Version=$(VERSION)" 
.PHONY: image
image:
ifdef IMAGE_ARCH
	docker buildx build --load --platform linux/$(IMAGE_ARCH) -t $(IMAGE_PREFIX)daikin-aircon-exporter:devel-$(IMAGE_ARCH) --build-arg VERSION=$(VERSION) .
else
	docker build -t $(IMAGE_PREFIX)daikin-aircon-exporter:devel --build-arg VERSION=$(VERSION) .
endif
.PHONY: push
push:
ifdef IMAGE_ARCH
	docker push $(IMAGE_PREFIX)daikin-aircon-exporter:$(IMAGE_TAG)-$(IMAGE_ARCH)
else
	docker push $(IMAGE_PREFIX)daikin-aircon-exporter:$(IMAGE_TAG)
endif

.PHONY: tag
tag:
ifdef IMAGE_ARCH
	docker tag $(IMAGE_PREFIX)daikin-aircon-exporter:devel-$(IMAGE_ARCH) $(IMAGE_PREFIX)daikin-aircon-exporter:$(IMAGE_TAG)-$(IMAGE_ARCH)
else
	docker tag $(IMAGE_PREFIX)daikin-aircon-exporter:devel $(IMAGE_PREFIX)daikin-aircon-exporter:$(IMAGE_TAG)
endif
.PHONY: test
test:
	test -z "$$(gofmt -s -l . | grep -v '^vendor' | tee /dev/stderr)"
	staticcheck ./...
	test -z "$$(nilerr ./... 2>&1 | tee /dev/stderr)"
	ineffassign .
	go install ./...
	go test -race -v ./...
	go vet ./...
	test -z "$$(go vet ./... | grep -v '^vendor' | tee /dev/stderr)"
.PHONY: tools
tools:
	cd /tmp; env GOFLAGS= GO111MODULE=on go get golang.org/x/tools/cmd/goimports
	cd /tmp; env GOFLAGS= GO111MODULE=on go get honnef.co/go/tools/cmd/staticcheck
	cd /tmp; env GOFLAGS= GO111MODULE=on go get github.com/gordonklaus/ineffassign
	cd /tmp; env GOFLAGS= GO111MODULE=on go get github.com/gostaticanalysis/nilerr/cmd/nilerr
