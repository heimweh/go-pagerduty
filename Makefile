GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=go-pagerduty

default: build

build:
	go install github.com/heimweh/go-pagerduty

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
	
vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)	xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

.PHONY: build test vet fmt
