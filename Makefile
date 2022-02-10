GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=go-pagerduty
FILES ?= "./..."
GOPKGS ?= $(shell go list $(FILES) | grep -v /vendor/)

default: build

build:
	@go get github.com/nordcloud/go-pagerduty/pagerduty

test:
	@echo "==> Testing ${PKG_NAME}"
	@go test -count 1 -timeout=30s -parallel=4 ${GOPKGS} ${TESTARGS}

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	@gofmt -w $(GOFMT_FILES)

.PHONY: build test vet fmt
