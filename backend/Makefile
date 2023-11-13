.PHONY: build test generate generate-wire

REPO_ROOT			:= $(shell git rev-parse --show-top-level)
PROJECT_ROOT		:= $(abspath $(dir $(MAKEFILE_LIST)))
PKG					= github.com/buildbeaver/buildbeaver
VERSION_INFO		= $(shell ../build/scripts/version-info.sh)
GIT_SHA_SHORT		= $(shell ../build/scripts/version-info.sh sha-short)
VERSION_VAR			=-X $(PKG)/common/version.VERSION=$(VERSION_INFO) -X $(PKG)/common/version.GITCOMMIT=$(GIT_SHA_SHORT)
GO_LDFLAGS			=-ldflags "$(VERSION_VAR)"

generate: generate-wire

generate-wire:
	@echo "--- generate wire ---"
	cd bb/app && wire

build:
	@echo "--- build ---"
	cd bb/cmd/bb && go install ${GO_LDFLAGS}

test:
	@echo "--- test ---"
	go test -mod=vendor ./...
