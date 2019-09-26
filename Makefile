# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

APIS := audit/v1 test/v1 tcc/core/v1 tcc/workflows/loadbalancer/v1

PROTOC_VERSION := 3.6.1
PROTOC_GEN_GO_VERSION := v1.2.0

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

PROTOC_INSTALLED := $(shell protoc --version | cut -d " " -f 2)
PROTOC_GEN_GO_INSTALLED := $(shell git -C $(GOSRCDIR)/github.com/golang/protobuf branch 2> \
	/dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/ \1/' | rev | cut -d " " -f 1 | cut -c2- | rev)

ifneq ($(PROTOC_INSTALLED),$(PROTOC_VERSION))
$(error Bad protoc version $(PROTOC_INSTALLED), please install $(PROTOC_VERSION))
endif

ifneq ($(PROTOC_GEN_GO_INSTALLED),$(PROTOC_GEN_GO_VERSION))
$(error Bad protoc-gen-go version $(PROTOC_GEN_GO_INSTALLED), \
	please install $(PROTOC_GEN_GO_VERSION))
endif

.PHONY: all
all: $(APIS)

.PHONY: $(APIS)
$(APIS):
	@echo "--- $@: all ---"
	$(MAKE) all -C $@

.PHONY: clean
clean:
	@for API in $(APIS); do \
		echo "--- $$API: $@ ---"; \
		$(MAKE) $@ -C $$API; \
		if [ $$? -ne 0 ]; then \
			exit 1; \
		fi; \
	done
