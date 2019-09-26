# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
DESCRIPTOR := test.proto-descriptor

.PHONY: all
all: $(DESCRIPTOR)
	@echo "All protos done"

${DESCRIPTOR}:
	protoc \
		-I $(TOPDIR)/ \
		-I $(VENDOR) \
		--go_out=plugins=grpc:. \
		--descriptor_set_out=test.proto-descriptor --include_imports \
		$(PROTO_SOURCES)

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR)
	@echo "Cleaned ${PROTO_GO_SOURCES} $(DESCRIPTOR)"
