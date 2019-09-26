# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

DESCRIPTOR = inventory.proto-descriptor

.PHONY: all
all: $(PROTO_GO_SOURCES)
	@echo "All protos done"

$(DESCRIPTOR): $(PROTO_GO_SOURCES)
$(PROTO_GO_SOURCES): $(PROTO_SOURCES)
	@protoc -I ./ \
		-I$(VENDOR) \
		--go_out=plugins=grpc:$(GOPATH)/src/ \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$?

clean:
	@rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR)
	@echo "Cleaned $(PROTO_GO_SOURCES) $(DESCRIPTOR)"
