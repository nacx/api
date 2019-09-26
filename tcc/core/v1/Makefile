# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../../..
VENDOR := $(TOPDIR)/third_party/protobuf/
APIS := $(TOPDIR)/api

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
PROTO_GO_VALIDATE_SOURCES := $(patsubst %.proto,%.pb.validate.go,$(PROTO_SOURCES))
PROTO_GO_GW_SOURCES = $(patsubst %.proto,%.pb.gw.go,$(PROTO_SOURCES))

DESCRIPTOR := core.proto-descriptor
SWAGGER := core.swagger.json

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

.PHONY: all
all: $(DESCRIPTOR)
	@echo "all protos done"

${DESCRIPTOR}:
	@echo "building protos"
	protoc -I ./ \
		-I $(VENDOR) \
		-I $(APIS) \
		-I $(GOSRCDIR)/github.com/envoyproxy/protoc-gen-validate \
		--swagger_out=json_names_for_fields=true,allow_merge=true,merge_file_name=core,fqn_for_swagger_name=true:. \
		--grpc-gateway_out=paths=source_relative:. \
		--go_out=paths=source_relative,plugins=grpc:. \
		--validate_out=paths=source_relative,lang=go:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$(PROTO_SOURCES)
	@echo "done building protos"

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(PROTO_GO_VALIDATE_SOURCES) $(PROTO_GO_GW_SOURCES) $(SWAGGER)
	@echo "cleaned ${PROTO_GO_SOURCES} $(PROTO_GO_VALIDATE_SOURCES) $(PROTO_GO_GW_SOURCES) $(SWAGGER) $(DESCRIPTOR)"
