# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

TOPDIR := ../../../..

include $(TOPDIR)/api/gogo.mk

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

DESCRIPTOR := spm-config.proto-descriptor

.PHONY: all
all: $(DESCRIPTOR)
	@echo "all protos done"

${DESCRIPTOR}:
	@echo "building protos"
	protoc -I ./ \
		-I /usr/include/protobuf \
		--gogofast_out=paths=source_relative,plugins=grpc,$(MAPS):. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		--validate_out=paths=source_relative,lang=gogo:. \
		$(PROTO_SOURCES)
	@echo "done building protos"

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR)
	@echo "cleaned ${PROTO_GO_SOURCES} $(DESCRIPTOR)"
