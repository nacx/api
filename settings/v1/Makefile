# Copyright (c) Tetrate, Inc 2020 All Rights Reserved.

TOPDIR := ../../..

API_NAME := settings

VENDOR := $(TOPDIR)/third_party/protobuf/
APIS := $(TOPDIR)/api

include $(TOPDIR)/api/gogo.mk

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
PROTO_GO_VALIDATE_SOURCES := $(patsubst %.proto,%.pb.validate.go,$(PROTO_SOURCES))
PROTO_GO_GW_SOURCES = $(patsubst %.proto,%.pb.gw.go,$(PROTO_SOURCES))

DESCRIPTOR := $(API_NAME).proto-descriptor
SWAGGER := $(API_NAME).swagger.json

all: $(DESCRIPTOR) $(SWAGGER)
	@echo "all protos done"

${DESCRIPTOR}:
	@echo "building $(API_NAME) protos"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		-I $(APIS) \
		--gogofast_out=paths=source_relative,plugins=grpc,$(MAPS):. \
		--grpc-gateway_out=allow_delete_body=true,paths=source_relative,$(MAPS):. \
		--validate_out=paths=source_relative,lang=gogo:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$(PROTO_SOURCES)
	@echo "done building $(API_NAME) protos"

# Generates swagger separately since we need to go on templating mechanism to potentially
# fill up host and base path differently between deployment.
$(SWAGGER):
	@echo "building $(API_NAME) swagger"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		-I $(APIS) \
		--swagger_out=allow_delete_body=true,json_names_for_fields=true,allow_merge=true,merge_file_name=$(API_NAME),fqn_for_swagger_name=true:. \
		$(PROTO_SOURCES)
	sed -i -e 's/\\n/ /g' $@
	@echo "done building $(API_NAME) swagger"

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(PROTO_GO_VALIDATE_SOURCES) $(PROTO_GO_GW_SOURCES) $(SWAGGER)
	@echo "cleaned ${PROTO_GO_SOURCES} $(PROTO_GO_VALIDATE_SOURCES) $(PROTO_GO_GW_SOURCES) $(SWAGGER) $(DESCRIPTOR)"

.PHONY: all clean
