# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

PROTO_SOURCES := $(wildcard *.proto)
PROTO_GO_SOURCES := $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
PROTO_GO_VALIDATE_SOURCES := $(patsubst %.proto,%.pb.validate.go,$(PROTO_SOURCES))

DESCRIPTOR := inventory.proto-descriptor
ENVOYFILTER := envoyfilter.yaml

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

.PHONY: all
all: $(DESCRIPTOR) $(ENVOYFILTER)
	@echo "all protos done"

$(DESCRIPTOR):
	@echo "building protos"
	GO111MODULE=off go get -d github.com/lyft/protoc-gen-validate
	GO111MODULE=off make -C $(GOSRCDIR)/github.com/lyft/protoc-gen-validate build
	protoc -I ./ \
		-I $(VENDOR) \
		-I $(GOSRCDIR)/github.com/lyft/protoc-gen-validate \
		--go_out=paths=source_relative,plugins=grpc:. \
		--descriptor_set_out=$(DESCRIPTOR) --include_imports \
		--validate_out=paths=source_relative,lang=go:. \
		$(PROTO_SOURCES)
	@echo "done building protos"

$(ENVOYFILTER): $(DESCRIPTOR)
	@echo "generating envoy filter..."
	GO111MODULE=off go get github.com/tetratelabs/istio-tools/grpc-transcoder
	grpc-transcoder \
		--descriptor $(DESCRIPTOR) \
		--service inventory-server \
		--port 9080 >$(ENVOYFILTER)

clean:
	rm -f $(PROTO_GO_SOURCES) $(PROTO_GO_VALIDATE_SOURCES) $(DESCRIPTOR) $(ENVOYFILTER)
	@echo "cleaned $(PROTO_GO_SOURCES) $(DESCRIPTOR) $(ENVOYFILTER)"
