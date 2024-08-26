#!/usr/bin/make -f

###############################################################################
###                               Go Modules                                ###
###############################################################################

verify:
	@echo "Verifying all go module dependencies..."
	@find . -name 'go.mod' -type f -execdir go mod verify \;

tidy:
	@echo "Cleaning up all go module dependencies..."
	@find . -name 'go.mod' -type f -execdir go mod tidy \;

.PHONY: verify tidy

###############################################################################
###                             Lint / Format                               ###
###############################################################################

lint:
	@echo "Linting all go modules..."
	@find . -name 'go.mod' -type f -execdir golangci-lint run --out-format=tab \;

lint-fix: format
	@echo "Attempting to fix lint errors in all go modules..."
	@find . -name 'go.mod' -type f -execdir golangci-lint run --fix --out-format=tab --issues-exit-code=0 \;

format_filter = \
	-type f \
	-name '*.go' \
	! -path '*/mocks/*' \
	! -name '*.cosmos_orm.go' \
	! -name '*.pb.go' \
	! -name '*.pb.gw.go' \
	! -name '*.pulsar.go' \
	! -name 'statik.go'

format:
	@echo "Formatting all go modules..."
	@find . $(format_filter) | xargs gofmt -s -w
	@find . $(format_filter) | xargs goimports -w -local github.com/chora-io/mods
	@find . $(format_filter) | xargs misspell -w

.PHONY: lint lint-fix format

###############################################################################
###                               Go Version                                ###
###############################################################################

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
MIN_GO_MAJOR_VERSION = 1
MIN_GO_MINOR_VERSION = 22
GO_VERSION_ERROR = Golang version $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION) is not supported, \
please update to at least $(MIN_GO_MAJOR_VERSION).$(MIN_GO_MINOR_VERSION)

go-version:
	@echo "Verifying go version..."
	@if [ $(GO_MAJOR_VERSION) -gt $(MIN_GO_MAJOR_VERSION) ]; then \
		exit 0; \
	elif [ $(GO_MAJOR_VERSION) -lt $(MIN_GO_MAJOR_VERSION) ]; then \
		echo $(GO_VERSION_ERROR); \
		exit 1; \
	elif [ $(GO_MINOR_VERSION) -lt $(MIN_GO_MINOR_VERSION) ]; then \
		echo $(GO_VERSION_ERROR); \
		exit 1; \
	fi

.PHONY: go-version

###############################################################################
###                                  Tools                                  ###
###############################################################################

tools: go-version
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: tools

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVersion=0.14.0
protoImage=ghcr.io/cosmos/proto-builder:$(protoVersion)
containerProtoFmt=chora-mods-proto-fmt-$(protoVersion)
containerProtoGen=chora-mods-proto-gen-$(protoVersion)

proto-all: proto-lint-fix proto-format proto-gen-admin proto-gen-agent proto-gen-example proto-gen-governor proto-gen-validator proto-gen-voucher proto-check-breaking

proto-lint:
	@protolint .

proto-lint-fix:
	@protolint -fix .

proto-format:
	@echo "Formatting protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find .  -name "*.proto" -exec clang-format -i {} \; ; fi

proto-gen-admin:
	@echo "Generating protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}-admin$$"; then docker start -a $(containerProtoGen)-admin; else docker run --name $(containerProtoGen)-admin -v $(CURDIR):/workspace --workdir /workspace $(protoImage) \
		sh -c 'cd admin; ./scripts/protocgen.sh'; fi

proto-gen-agent:
	@echo "Generating protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}-agent$$"; then docker start -a $(containerProtoGen)-agent; else docker run --name $(containerProtoGen)-agent -v $(CURDIR):/workspace --workdir /workspace $(protoImage) \
		sh -c 'cd agent; ./scripts/protocgen.sh'; fi

proto-gen-example:
	@echo "Generating protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}-example$$"; then docker start -a $(containerProtoGen)-example; else docker run --name $(containerProtoGen)-example -v $(CURDIR):/workspace --workdir /workspace $(protoImage) \
		sh -c 'cd example; ./scripts/protocgen.sh'; fi

proto-gen-governor:
	@echo "Generating protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}-governor$$"; then docker start -a $(containerProtoGen)-governor; else docker run --name $(containerProtoGen)-governor -v $(CURDIR):/workspace --workdir /workspace $(protoImage) \
		sh -c 'cd governor; ./scripts/protocgen.sh'; fi

proto-gen-validator:
	@echo "Generating protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}-validator$$"; then docker start -a $(containerProtoGen)-validator; else docker run --name $(containerProtoGen)-validator -v $(CURDIR):/workspace --workdir /workspace $(protoImage) \
		sh -c 'cd validator; ./scripts/protocgen.sh'; fi

proto-gen-voucher:
	@echo "Generating protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}-voucher$$"; then docker start -a $(containerProtoGen)-voucher; else docker run --name $(containerProtoGen)-voucher -v $(CURDIR):/workspace --workdir /workspace $(protoImage) \
		sh -c 'cd voucher; ./scripts/protocgen.sh'; fi

proto-check-breaking:
	@docker run -v $(shell pwd):/workspace --workdir /workspace bufbuild/buf:1.9.0 breaking --against https://github.com/chora-io/mods.git#branch=main

.PHONY: proto-all proto-lint proto-lint-fix proto-format proto-gen-admin proto-gen-agent proto-gen-example proto-gen-governor proto-gen-validator proto-gen-voucher proto-check-breaking

###############################################################################
###                                  Tests                                  ###
###############################################################################

CURRENT_DIR=$(shell pwd)
GO_MODULES=$(shell find . -type f -name 'go.mod' -print0 | xargs -0 -n1 dirname | sort)

test: test-all

test-all:
	@for module in $(GO_MODULES); do \
		echo "Testing Module $$module"; \
		cd ${CURRENT_DIR}/$$module; \
		go test ./...; \
	done

test-admin:
	@echo "Testing Module admin"
	@cd admin && go test ./... \
		-coverprofile=../coverage-admin.out -covermode=atomic

test-agent:
	@echo "Testing Module agent"
	@cd agent && go test ./... \
		-coverprofile=../coverage-agent.out -covermode=atomic

test-example:
	@echo "Testing Module example"
	@cd example && go test ./... \
		-coverprofile=../coverage-example.out -covermode=atomic

test-governor:
	@echo "Testing Module governor"
	@cd governor && go test ./... \
		-coverprofile=../coverage-governor.out -covermode=atomic

test-validator:
	@echo "Testing Module validator"
	@cd validator && go test ./... \
		-coverprofile=../coverage-validator.out -covermode=atomic

test-voucher:
	@echo "Testing Module voucher"
	@cd voucher && go test ./... \
		-coverprofile=../coverage-voucher.out -covermode=atomic

test-coverage:
	@cat coverage*.out | grep -v "mode: atomic" >> coverage.txt

test-clean:
	@go clean -testcache
	@find . -name 'coverage.txt' -delete
	@find . -name 'coverage*.out' -delete

.PHONY: test test-all test-admin test-agent test-example test-governor test-validator test-voucher test-coverage test-clean

###############################################################################
###                              Documentation                              ###
###############################################################################

docs:
	@echo "Wait a few seconds and then visit http://localhost:6060/pkg/github.com/chora-io/mods/"
	godoc -http=:6060

.PHONY: docs

###############################################################################
###                                 Clean                                   ###
###############################################################################

clean: test-clean
	@rm -rf $(BUILD_DIR)

.PHONY: clean
