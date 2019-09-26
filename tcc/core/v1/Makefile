# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.
NPM := $(or ${shell which npm},npm)
YARN := $(or ${shell which yarn},yarn)
NODE := $(or ${shell which node},node)
OPENAPI_GEN := $(or ${shell which openapi-generator},openapi-generator)

TOPDIR := ../../../..
VENDOR := $(TOPDIR)/third_party/protobuf/
APIS := $(TOPDIR)/api
TS_OUTPUT_GEN := ./generated/ts

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
PROTO_GO_VALIDATE_SOURCES := $(patsubst %.proto,%.pb.validate.go,$(PROTO_SOURCES))
PROTO_GO_GW_SOURCES = $(patsubst %.proto,%.pb.gw.go,$(PROTO_SOURCES))

TS_APIS = ./generated/ts/apis/*
TS_MODELS = ./generated/ts/models/*
TS_INDEX = ./generated/ts/index.ts
TS_RUNTIME = ./generated/ts/runtime.ts

DESCRIPTOR := core.proto-descriptor
SWAGGER := core.swagger.json

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

.PHONY: all
# TODO(aji): generate-ts isn't included here, because it breaks in CI.
# To not break `make check` in CI, we need to have JVM running.
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
	# TODO(aji): uncomment when generate-ts is included in `make all`.
	# rm -f $(TS_APIS) $(TS_MODELS) $(TS_INDEX) $(TS_RUNTIME)
	# @echo "cleaned $(TS_APIS) $(TS_MODELS) $(TS_INDEX) $(TS_RUNTIME)"

generate-ts: $(OPENAPI_GEN)
	@echo "generating typescript-fetch from swagger"
	$(OPENAPI_GEN) generate -i ${SWAGGER} -g typescript-fetch -o ${TS_OUTPUT_GEN}
	@echo "done generating typescript files"

$(NPM):
	$(error npm is required)

$(NODE):
	$(error node is required)

$(YARN): $(NPM)
	@echo installing yarn
	@echo npm install -g yarn

$(OPENAPI_GEN): $(YARN)
	@echo installing openapi-generator
	yarn global add @openapitools/openapi-generator-cli@0.0.9-4.0.0-beta3
