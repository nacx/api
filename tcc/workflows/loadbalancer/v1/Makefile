# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.
TOPDIR := ../../../../..

NPM := $(or ${shell which npm},npm)
YARN := $(TOPDIR)/scripts/yarn
NODE := $(or ${shell which node},node)
OPENAPI_GEN := $(or ${shell which openapi-generator},openapi-generator)

VENDOR := $(TOPDIR)/third_party/protobuf/
APIS := $(TOPDIR)/api
TS_OUTPUT_GEN := ./generated/ts

include $(TOPDIR)/api/gogo.mk

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
PROTO_GO_VALIDATE_SOURCES := $(patsubst %.proto,%.pb.validate.go,$(PROTO_SOURCES))
PROTO_GO_GW_SOURCES = $(patsubst %.proto,%.pb.gw.go,$(PROTO_SOURCES))

TS_APIS = ./generated/ts/apis/*
TS_MODELS = ./generated/ts/models/*
TS_INDEX = ./generated/ts/index.ts
TS_RUNTIME = ./generated/ts/runtime.ts
TS_GENERATED = ./generated/ts/tsconfig.json

DESCRIPTOR := workflows.proto-descriptor
SWAGGER := workflows.swagger.json

all: $(DESCRIPTOR) $(SWAGGER) $(TS_GENERATED)
	@echo "all protos done"

${DESCRIPTOR}:
	@echo "building protos"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		-I $(APIS) \
		--gogofast_out=paths=source_relative,plugins=grpc,$(MAPS):. \
		--grpc-gateway_out=allow_delete_body=true,paths=source_relative,$(MAPS):. \
		--validate_out=paths=source_relative,lang=gogo:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$(PROTO_SOURCES)
	@echo "done building protos"

# Generates swagger separately since we need to go on templating mechanism to potentially
# fill up host and base path differently between deployment.
$(SWAGGER):
	@echo "building swagger"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		-I $(APIS) \
		--swagger_out=allow_delete_body=true,json_names_for_fields=true,allow_merge=true,merge_file_name=workflows,fqn_for_swagger_name=true:. \
		$(PROTO_SOURCES)
	sed -i -e 's/\\n/ /g' $@
	rm -fr $(TS_OUTPUT_GEN)
	@echo "done building swagger"

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(PROTO_GO_VALIDATE_SOURCES) $(PROTO_GO_GW_SOURCES) $(SWAGGER)
	@echo "cleaned ${PROTO_GO_SOURCES} $(PROTO_GO_VALIDATE_SOURCES) $(PROTO_GO_GW_SOURCES) $(SWAGGER) $(DESCRIPTOR)"
	rm -f $(TS_APIS) $(TS_MODELS) $(TS_INDEX) $(TS_RUNTIME)
	@echo "cleaned $(TS_APIS) $(TS_MODELS) $(TS_INDEX) $(TS_RUNTIME)"

$(TS_GENERATED): $(OPENAPI_GEN) $(YARN) $(NPM) $(NODE) $(SWAGGER)
	@echo "generating typescript-fetch from swagger"
	$(OPENAPI_GEN) generate -i ${SWAGGER} -g typescript-fetch -o ${TS_OUTPUT_GEN}
	@echo "done generating typescript files"

$(NPM):
	$(error npm is required)

$(NODE):
	$(error node is required)

$(YARN):
	wget -O - -q https://github.com/yarnpkg/yarn/releases/download/v1.17.3/yarn-1.17.3.js > $(YARN)

$(OPENAPI_GEN): $(YARN)
	$(YARN) global add gen8

.PHONY: all clean
