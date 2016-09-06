APPLICATION := $$(basename $(CURDIR))
BUILD_DIR ?= bin

NOROOT := -u $$(id -u):$$(id -g)
SRCDIR := /go/src/github.com/gravitational/$(APPLICATION)
DOCKERFLAGS := --rm=true $(NOROOT) -v $(PWD):$(SRCDIR) -w $(SRCDIR)
BUILDIMAGE := quay.io/gravitational/debian-venti:go1.7-jessie

.PHONY: all
all: clean build

$(BUILD_DIR):
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o $(BUILD_DIR)/$(APPLICATION) .

.PHONY: build
build:
	docker run $(DOCKERFLAGS) $(BUILDIMAGE) make $(BUILD_DIR)

.PHONY: test
test:
	go vet .

.PHONY: clean
clean:
	-rm -r $(BUILD_DIR)

.PHONY: install-tools
install-tools:
	go get -u github.com/mattn/goreman
