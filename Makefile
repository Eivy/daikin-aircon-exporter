.PHONY: build/daikin-aircon-exporter
build/daikin-aircon-exporter:
	mkdir -p build
	go build -o $@ -ldflags "-X github.com/eivy/daikin-aircon-exporter.Version=$(VERSION)" 
.PHONY: image
image:
	docker buildx build --load --platform linux/arm64 -t $(IMAGE_PREFIX)daikin-aircon-exporter:devel --build-arg VERSION=$(VERSION) .
.PHONY: tag
tag:
	docker tag $(IMAGE_PREFIX)topolvm:devel $(IMAGE_PREFIX)topolvm:$(IMAGE_TAG)
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
