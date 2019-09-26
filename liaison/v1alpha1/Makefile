# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

ISTIO_API_VERSION := 1.1.0-snapshot.5

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

TOPDIR := ../../..
VENDOR := $(TOPDIR)/third_party/protobuf/

TOP_PROTOS := $(wildcard *.proto)
ALL_PROTOS := $(wildcard *.proto) $(wildcard */*.proto)
ALL_GO_PROTOS := $(patsubst %.proto,%.pb.go,$(ALL_PROTOS))
ALL_GO_VALIDATE_PROTOS := $(patsubst %.proto,%.pb.validate.go,$(ALL_PROTOS))

ALL_JS_PROTOS := $(patsubst %.proto,%_pb.js,$(ALL_PROTOS))
ALL_JS_SERVICE_PROTOS := $(patsubst %.proto,%_pb_service.js,$(ALL_PROTOS))
ALL_TS_PROTOS := $(patsubst %.proto,%_pb.d.ts,$(ALL_PROTOS))
ALL_TS_SERVICE_PROTOS := $(patsubst %.proto,%_pb_service.d.ts,$(ALL_PROTOS))
ALL_CLIENT_PROTOS := $(ALL_JS_PROTOS) $(ALL_TS_PROTOS) $(ALL_JS_SERVICE_PROTOS) \
	$(ALL_TS_SERVICE_PROTOS)

ALL_JS_ARTIFACTS := ./google ./validate node_modules

# Istio API.
# TODO(dio): We should have a better naming for ISTIO_API if tar in osx supports --transform.
ISTIO_API := api-$(ISTIO_API_VERSION)

# Targets.
DESCRIPTOR := liaison.proto-descriptor
ENVOYFILTER := envoyfilter.yaml

ALL_GENERATED := $(ALL_GO_PROTOS) $(ALL_GO_VALIDATE_PROTOS) $(DESCRIPTOR) $(ENVOYFILTER) \
	$(ISTIO_API) $(ALL_CLIENT_PROTOS) $(ALL_JS_ARTIFACTS)

.PHONY: build
build: $(DESCRIPTOR) $(ENVOYFILTER) node_modules
	@echo "all protos done"

# Install all the required global deps.
define install_deps
	make $(ISTIO_API)
	@yarn
endef

# Execute protoc gen go and generate the .proto-descriptor.
define protoc_gen_go
	$(call install_deps)
	protoc -I ./ \
		-I $(VENDOR) \
		-I $(ISTIO_API) \
		-I $(GOSRCDIR)/github.com/lyft/protoc-gen-validate \
		--go_out=paths=source_relative,plugins=grpc:. \
		--validate_out=paths=source_relative,lang=go:. \
		$(2) \
		$(3)
endef

# Execute protoc gen ts (installed in ./node_modules/.bin/protoc-gen-ts).
define protoc_gen_ts
	$(call install_deps)
	protoc -I ./ \
		-I $(VENDOR) \
		-I $(ISTIO_API) \
		-I $(GOSRCDIR)/github.com/lyft/protoc-gen-validate \
		--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
		--js_out=import_style=commonjs,binary:. \
  		--ts_out=service=true:. \
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

$(ALL_GO_PROTOS):
	$(call protoc_gen_go,$@,``,$(patsubst %.pb.go,%.proto,$@))

$(DESCRIPTOR): $(ALL_GO_PROTOS)
	$(call protoc_gen_go,$@,--descriptor_set_out=$(DESCRIPTOR) --include_imports,$(TOP_PROTOS))

$(ENVOYFILTER): $(DESCRIPTOR)
	@echo "generating envoy filter"
	$(call grpc_transcoder,$@,9080)

$(ALL_JS_PROTOS):
	$(call protoc_gen_ts,$@,``,$(patsubst %_pb.js,%.proto,$@))

$(ALL_TS_PROTOS):
	$(call protoc_gen_ts,$@,``,$(patsubst %_pb.d.ts,%.proto,$@))

$(ISTIO_API):
	@wget -q -O- https://github.com/istio/api/archive/$(ISTIO_API_VERSION).tar.gz \
	| tar xz $(ISTIO_API)/networking/v1alpha3

node_modules: touch
	@yarn

.PHONY: js touch
js: node_modules $(ALL_JS_PROTOS) $(ALL_TS_PROTOS)

touch:
	mkdir -p ./google/api
	mkdir -p ./validate
	touch ./google/api/annotations_pb.js
	touch ./validate/validate_pb.js

clean:
	rm -fr $(ALL_GENERATED)
	@echo "cleaned $(ALL_GENERATED)"
