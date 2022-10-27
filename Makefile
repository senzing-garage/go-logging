# Makefile that builds a "go" program.

# "Simple expanded" variables (':=')

# PROGRAM_NAME is the name of the GIT repository.
PROGRAM_NAME := $(shell basename `git rev-parse --show-toplevel`)
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MAKEFILE_DIRECTORY := $(dir $(MAKEFILE_PATH))
TARGET_DIRECTORY := $(MAKEFILE_DIRECTORY)/target
BUILD_VERSION := $(shell git describe --always --tags --abbrev=0 --dirty)
BUILD_TAG := $(shell git describe --always --tags --abbrev=0)
BUILD_ITERATION := $(shell git log $(BUILD_TAG)..HEAD --oneline | wc -l | sed 's/^ *//')
GIT_REMOTE_URL := $(shell git config --get remote.origin.url)
GO_PACKAGE_NAME := $(shell echo $(GIT_REMOTE_URL) | sed -e 's|^git@github.com:|github.com/|' -e 's|\.git$$||')

# The first "make" target runs as default.

.PHONY: default
default: help

# ---- Linux ------------------------------------------------------------------

target/linux:
	@mkdir -p $(TARGET_DIRECTORY)/linux || true


target/linux/$(PROGRAM_NAME): target/linux
	GOOS=linux \
	GOARCH=amd64 \
	go build \
		-a \
		-ldflags " \
			-X main.programName=${PROGRAM_NAME} \
			-X main.buildVersion=${BUILD_VERSION} \
			-X main.buildIteration=${BUILD_ITERATION} \
			" \
		-o $(TARGET_DIRECTORY)/linux/$(PROGRAM_NAME)


# -----------------------------------------------------------------------------
# Build
#   Notes:
#     "-a" needed to incorporate changes to C files.
# -----------------------------------------------------------------------------

.PHONY: dependencies
dependencies:
	@go get -u
	@go get ./...
	@go mod tidy
	@go get -u github.com/jstemmer/go-junit-report


.PHONY: build
build: dependencies \
	target/linux/$(PROGRAM_NAME)


# -----------------------------------------------------------------------------
# Test
# -----------------------------------------------------------------------------

.PHONY: test
test:
	@go test -v ./...
#	@go test -v ./.
#	@go test -v ./logger
#	@go test -v ./messageformat
#	@go test -v ./messageid
#	@go test -v ./messagelogger
#	@go test -v ./messageloglevel
#	@go test -v ./messagestatus
#	@go test -v ./messagetext


# -----------------------------------------------------------------------------
# Run
# -----------------------------------------------------------------------------

.PHONY: run
run:
	@target/linux/$(PROGRAM_NAME)

# -----------------------------------------------------------------------------
# Utility targets
# -----------------------------------------------------------------------------

.PHONY: clean
clean:
	@go clean -cache
	@rm -rf $(TARGET_DIRECTORY) || true
	@rm -f $(GOPATH)/bin/$(PROGRAM_NAME) || true


.PHONY: print-make-variables
print-make-variables:
	@$(foreach V,$(sort $(.VARIABLES)), \
		$(if $(filter-out environment% default automatic, \
		$(origin $V)),$(warning $V=$($V) ($(value $V)))))


.PHONY: help
help:
	@echo "Build $(PROGRAM_NAME) version $(BUILD_VERSION)-$(BUILD_ITERATION)".
	@echo "All targets:"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs
