# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../../..

VENDOR := $(TOPDIR)/third_party/protobuf/

include $(TOPDIR)/api/gogo.mk

PROTO_SOURCES := $(wildcard *.proto)
PROTO_GO_SOURCES := $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
PROTO_GO_VALIDATE_SOURCES := $(patsubst %.proto,%.pb.validate.go,$(PROTO_SOURCES))
PROTO_GO_GW_SOURCES = $(patsubst %.proto,%.pb.gw.go,$(PROTO_SOURCES))

DESCRIPTOR := spm-notification.proto-descriptor
SWAGGER := spm.swagger.json

.PHONY: all
all: $(DESCRIPTOR) $(SWAGGER)
	@echo "all protos done"

$(DESCRIPTOR):
	@echo "building protos"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		--grpc-gateway_out=allow_delete_body=true,paths=source_relative,$(MAPS):. \
		--gogofast_out=paths=source_relative,plugins=grpc,$(MAPS):. \
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
		--swagger_out=allow_delete_body=true,json_names_for_fields=true,allow_merge=true,merge_file_name=spm,fqn_for_swagger_name=true:. \
		$(PROTO_SOURCES)
	sed -i -e 's/\\n/ /g' $@
	@echo "done building swagger"

.PHONY: clean
clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(PROTO_GO_VALIDATE_SOURCES) $(PROTO_GO_GW_SOURCES) $(SWAGGER)
	@echo "cleaned $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(PROTO_GO_VALIDATE_SOURCES) $(PROTO_GO_GW_SOURCES) $(SWAGGER)"
