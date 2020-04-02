## Copyright (c) Tetrate, Inc 2020 All Rights Reserved.

TOPDIR := ../../../..

VENDOR := $(TOPDIR)/third_party/protobuf/
APIS := $(TOPDIR)/api

include $(TOPDIR)/api/gogo.mk

PROTO_SOURCES = $(wildcard *.proto)
PROTO_GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
DONE := .done

all: $(DONE)
	@echo "all protos done"

$(DONE):
	@echo "building protos"
	protoc -I ./ \
		-I /usr/include/protobuf \
		-I $(VENDOR) \
		-I $(APIS) \
		--go_out=paths=source_relative,plugins=grpc,$(MAPS):. \
		--docs_out=warnings=true,per_file=true,mode=html_fragment_with_front_matter:. \
		--include_imports $(PROTO_SOURCES)
	@echo "done building protos"
	@echo "running values_types.pb.go fixup structs"
	cd /tmp; go run $(CURDIR)/../../common/fixup_structs/main.go -f $(CURDIR)/values_types.pb.go
	@echo "running component.pb.go fixup structs"
	cd /tmp; go run $(CURDIR)/../../common/fixup_structs/main.go -f $(CURDIR)/component.pb.go
	@echo "running operator.pb.go fixup structs"
	@cd /tmp; go run $(CURDIR)/../../common/fixup_structs/main.go -f $(CURDIR)/operator.pb.go
	@touch .done


clean:
	rm -f $(PROTO_GO_SOURCES) $(DONE)
	@echo "cleaned ${PROTO_GO_SOURCES} $(DONE)"

.PHONY: all clean
