# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
PROTO_GO_VALIDATE_SOURCES := $(patsubst %.proto,%.pb.validate.go,$(PROTO_SOURCES))

DESCRIPTOR := test.proto-descriptor

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

.PHONY: all
all: $(DESCRIPTOR)
	@echo "all protos done"

${DESCRIPTOR}:
	@echo "building protos"
	GO111MODULE=off go get -d github.com/lyft/protoc-gen-validate
	GO111MODULE=off make -C $(GOSRCDIR)/github.com/lyft/protoc-gen-validate build
	protoc -I ./ \
		-I $(VENDOR) \
		-I $(GOSRCDIR)/github.com/lyft/protoc-gen-validate \
		--go_out=paths=source_relative,plugins=grpc:. \
		--validate_out=paths=source_relative,lang=go:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$(PROTO_SOURCES)
	@echo "done building protos"

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(PROTO_GO_VALIDATE_SOURCES)
	@echo "cleaned ${PROTO_GO_SOURCES} $(PROTO_GO_VALIDATE_SOURCES) $(DESCRIPTOR)"
