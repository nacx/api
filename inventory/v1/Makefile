# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

$(PROTO_GO_SOURCES): $(PROTO_SOURCES)
	@protoc -I ./ \
		-I$(VENDOR) \
		--go_out=plugins=grpc:$(GOPATH)/src/ \
		--descriptor_set_out=inventory.proto-descriptor --include_imports \
		$?

all: $(PROTO_GO_SOURCES)
	@echo "All protos done"

.PHONY: all

clean:
	@echo "Cleaned nothing"
