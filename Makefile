export GO111MODULE := on
export PATH := ./bin:$(PATH)

ci: bootstrap lint test
.PHONY: ci

#################################################
# Bootstrapping for base golang package and tool deps
#################################################

bootstrap:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.21.0
.PHONY: bootstrap

mod-update:
	go get -u -m
	go mod tidy

mod-tidy:
	go mod tidy

.PHONY: $(CMD_PKGS)
.PHONY: mod-update mod-tidy

#################################################
# Test and linting
#################################################
# Run all the linters
lint:
	bin/golangci-lint run ./...
.PHONY: lint

test:
	CGO_ENABLED=0 go test $$(go list ./... | grep -v generated)
.PHONY: test
