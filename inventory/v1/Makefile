# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ${PWD}
VENDOR := ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

${PROTO_GO_SOURCES}: ${PROTO_SOURCES}
	@protoc -I ${TOPDIR}/ \
		-I${VENDOR} \
		--go_out=plugins=grpc:${GOPATH}/src/ \
		--descriptor_set_out=inventory.proto-descriptor --include_imports \
		$?

all: ${PROTO_GO_SOURCES}
	@echo "All protos done"

.PHONY: all

clean:
	@echo "Cleaned nothing"
