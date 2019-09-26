# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

# Istio API.
ISTIO_API_VERSION := 1.1.0-snapshot.5
# TODO(dio): We should have a better naming for ISTIO_API if tar in osx supports --transform.
ISTIO_API := api-$(ISTIO_API_VERSION)
# The enabled Istio APIs.
ISTIO_APIS := $(ISTIO_API)/networking/v1alpha3 $(ISTIO_API)/mcp/v1alpha1

# Targets.
DESCRIPTOR := liaison.proto-descriptor
ENVOYFILTER := envoyfilter.yaml

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

ALL_PROTOS := $(wildcard *.proto)
ALL_GO_PROTOS := $(patsubst %.proto,%.pb.go,$(ALL_PROTOS))
ALL_GO_VALIDATE_PROTOS := $(patsubst %.proto,%.pb.validate.go,$(ALL_PROTOS))

ALL_GENERATED := $(ALL_GO_PROTOS) $(ALL_GO_VALIDATE_PROTOS) $(DESCRIPTOR) $(ENVOYFILTER) \
	$(ISTIO_API)

.PHONY: all
all: $(DESCRIPTOR) $(ENVOYFILTER)
	@echo "all protos done"

# Install all the required global deps.
define install_deps
	$(MAKE) $(ISTIO_API)
endef

# Execute protoc gen go and generate the .proto-descriptor.
define protoc_gen_go
	$(call install_deps)
	protoc -I ./ \
		-I $(VENDOR) \
		-I $(ISTIO_API) \
		-I $(GOSRCDIR)/github.com/envoyproxy/protoc-gen-validate \
		--go_out=paths=source_relative,plugins=grpc:. \
		--validate_out=paths=source_relative,lang=go:. \
		$(2) \
		$(3)
endef

# Execute grpc_transcoder to generate envoyfilter.yaml with a specified port.
define grpc_transcoder
	GO111MODULE=off go get github.com/tetratelabs/istio-tools/grpc-transcoder
	grpc-transcoder \
		--descriptor $(DESCRIPTOR) \
		--service liaison-server \
		--port $(2) > $(ENVOYFILTER)
endef

$(DESCRIPTOR):
	$(call protoc_gen_go,$@,--descriptor_set_out=$(DESCRIPTOR) --include_imports,$(ALL_PROTOS))

$(ENVOYFILTER): $(DESCRIPTOR)
	@echo "generating envoy filter"
	$(call grpc_transcoder,$@,9080)

$(ISTIO_API):
	@wget -q -O- https://github.com/istio/api/archive/$(ISTIO_API_VERSION).tar.gz \
	| tar xz $(ISTIO_APIS)

clean:
	rm -fr $(ALL_GENERATED)
	@echo "cleaned $(ALL_GENERATED)"
