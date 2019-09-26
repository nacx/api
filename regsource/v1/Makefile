# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

DESCRIPTOR := regsource.proto-descriptor

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

all: $(DESCRIPTOR) 
	@echo "all protos done"

${DESCRIPTOR}:
	@echo "building protos"
	protoc -I ./ \
		-I $(VENDOR) \
		--go_out=paths=source_relative,plugins=grpc:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$(PROTO_SOURCES)
	@echo "done building protos"


clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR) 
	@echo "cleaned ${PROTO_GO_SOURCES} $(DESCRIPTOR)"


.PHONY: all clean 
