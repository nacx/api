# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ${PWD}
VENDOR := ${GOPATH}/src/github.com/tetrateio/tetrate/third_party/protobuf

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

${PROTO_GO_SOURCES}: ${PROTO_SOURCES}
	@protoc \
		-I ${TOPDIR}/ \
		-I ${VENDOR} \
		--go_out=plugins=grpc:. \
		--descriptor_set_out=test.proto-descriptor --include_imports \
		$?

all: ${PROTO_GO_SOURCES}
	@echo "All protos done"

.PHONY: all

clean:
	@rm -f ${PROTO_GO_SOURCES} test.proto-descriptor
	@echo "Cleaned ${PROTO_GO_SOURCES}"
