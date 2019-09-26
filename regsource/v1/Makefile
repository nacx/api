# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

include $(TOPDIR)/api/gogo.mk

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

DESCRIPTOR := regsource.proto-descriptor

all: $(DESCRIPTOR)
	@echo "all protos done"

${DESCRIPTOR}:
	@echo "building protos"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		--gogofast_out=paths=source_relative,plugins=grpc,$(MAPS):. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$(PROTO_SOURCES)
	@echo "done building protos"

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR)
	@echo "cleaned ${PROTO_GO_SOURCES} $(DESCRIPTOR)"

.PHONY: all clean
