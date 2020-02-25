## Copyright (c) Tetrate, Inc 2020 All Rights Reserved.

TOPDIR := ../../..

VENDOR := $(TOPDIR)/third_party/protobuf/

include $(TOPDIR)/api/gogo.mk

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
DONE := .done
CURRDIR = $(shell pwd)

all: $(DONE)
	@echo "all protos done"

$(DONE):
	@echo "building protos"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		--go_out=paths=source_relative,plugins=grpc,$(MAPS):. \
		--include_imports $(PROTO_SOURCES)
	@echo "done building protos"
	@echo "running kubernetes.pb.go fixup structs"
	cd /tmp; go run $(CURRDIR)/fixup_structs/main.go -f $(CURRDIR)/kubernetes.pb.go
	@touch .done

clean:
	rm -f $(PROTO_GO_SOURCES) $(DONE)
	@echo "cleaned ${PROTO_GO_SOURCES} $(DONE)"

.PHONY: all clean
