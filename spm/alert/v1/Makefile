# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

TOPDIR := ../../../..

VENDOR := $(TOPDIR)/third_party/protobuf/
APIS := $(TOPDIR)/api
RBAC := $(APIS)/q/rbac/v1

include $(TOPDIR)/api/gogo.mk

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

DESCRIPTOR := spm-alert.proto-descriptor

.PHONY: all
all: $(DESCRIPTOR)
	@echo "all protos done"

${DESCRIPTOR}:
	@echo "building protos"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		-I $(APIS) \
		-I $(RBAC) \
		--gogofast_out=paths=source_relative,plugins=grpc,$(MAPS):. \
		--grpc-gateway_out=allow_delete_body=true,paths=source_relative,$(MAPS):. \
		--validate_out=paths=source_relative,lang=gogo:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		--docs_out=warnings=true,per_file=true,mode=html_fragment_with_front_matter:. \
		$(PROTO_SOURCES)
	@echo "done building protos"

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR)
	@echo "cleaned ${PROTO_GO_SOURCES} $(DESCRIPTOR)"
