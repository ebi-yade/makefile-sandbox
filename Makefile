GO := go
GO_BUILD := $(GO) build
GO_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
export GOBIN ?= $(shell pwd)/bin
export GO111MODULE := on
# export GOPRIVATE := github.com/ebi-yade/makefile-sandbox

DIST_DIR := _dist
# the value should be in the format of basename
LAMBDA_HANDLER_DIR := _handler
rec_wildcard = $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))
ENTIRE_SRC := $(call rwildcard,,*.go) go.mod go.sum
LAMBDA_HANDLERS := $(notdir $(patsubst %/,%,$(sort $(dir $(wildcard ./$(LAMBDA_HANDLER_DIR)/*/*.go)))))
# LAMBDA_HANDLERS_NOT_FRIENDLY_FOR_WINDOWS :=
# $(shell find $(LAMBDA_HANDLER_DIR)/* -maxdepth 1 -type d -execdir basename '{}' ';' 2>/dev/null)
TARGETS := $(LAMBDA_HANDLERS:%=$(DIST_DIR)/%)

$(DIST_DIR)/%: $(LAMBDA_HANDLER_DIR)/% $(ENTIRE_SRC)
	cd $< && $(GO_ENV) $(GO_BUILD) -ldflags="-s -w -X main.Revision=$(GIT_REV)" -o ../../$@

.DEFAULT_GOAL := handlers
.PHONY: handlers clean

handlers: $(TARGETS)

clean:
	$(GO) clean
	rm -rf $(DIST_DIR)
