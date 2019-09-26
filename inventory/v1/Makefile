# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

PROTO_SOURCES := $(wildcard *.proto)
PROTO_GO_SOURCES := $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

DESCRIPTOR := inventory.proto-descriptor
ENVOYFILTER := envoyfilter.yaml

.PHONY: all
all: $(DESCRIPTOR) $(ENVOYFILTER)
	@echo "All protos done"

$(DESCRIPTOR):
	@echo "building protos..."
	protoc -I ./ \
		-I$(VENDOR) \
		--go_out=paths=source_relative,plugins=grpc:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		$(PROTO_SOURCES)

$(ENVOYFILTER): $(DESCRIPTOR)
	@echo "generating Envoy filter..."
	GO111MODULE=off go get -u github.com/tetratelabs/istio-tools/grpc-transcoder
	GO111MODULE=off go run github.com/tetratelabs/istio-tools/grpc-transcoder \
		--descriptor $(DESCRIPTOR) \
		--service inventory-server \
		--port 9080 >$(ENVOYFILTER)

clean:
	rm -f $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(ENVOYFILTER)
	@echo "Cleaned $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(ENVOYFILTER)"
