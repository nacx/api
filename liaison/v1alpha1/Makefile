# Copyright (c) Tetrate, Inc 2018 All Rights Reserved.

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

DESCRIPTOR := liaison.proto-descriptor
ENVOYFILTER := envoyfilter.yaml

ALL_GENERATED := $(ALL_GO_PROTOS) $(ALL_GO_VALIDATE_PROTOS) $(DESCRIPTOR) $(ENVOYFILTER) \
	$(ALL_CLIENT_PROTOS)

GOPATH := $(shell go env GOPATH)
GOSRCDIR := $(GOPATH)/src

.PHONY: all
all: $(DESCRIPTOR) $(ENVOYFILTER) node_modules
	@echo "all protos done"

# Install all the required global deps.
define install_deps
	@GO111MODULE=off go get -d github.com/lyft/protoc-gen-validate
	@GO111MODULE=off make -C $(GOSRCDIR)/github.com/lyft/protoc-gen-validate build
	@GO111MODULE=off go get -d istio.io/api/networking/v1alpha3
	@yarn
endef

# Execute protoc gen go and generate the .proto-descriptor.
define protoc_gen_go
	$(call install_deps)
	protoc -I ./ \
		-I $(VENDOR) \
		-I $(GOSRCDIR)/github.com/lyft/protoc-gen-validate \
		-I $(GOSRCDIR)/istio.io/api \
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
		-I $(GOSRCDIR)/github.com/lyft/protoc-gen-validate \
		-I $(GOSRCDIR)/istio.io/api \
		--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
		--js_out=import_style=commonjs,binary:. \
  		--ts_out=service=true:. \
		$(2) \
		$(3)
endef

$(ALL_GO_PROTOS):
	$(call protoc_gen_go,$@,``,$(patsubst %.pb.go,%.proto,$@))

$(DESCRIPTOR): $(ALL_GO_PROTOS)
	$(call protoc_gen_go,$@,--descriptor_set_out=$(DESCRIPTOR) --include_imports,$(TOP_PROTOS))

$(ENVOYFILTER): $(DESCRIPTOR)
	@echo "generating envoy filter..."
	GO111MODULE=off go get github.com/tetratelabs/istio-tools/grpc-transcoder
	grpc-transcoder \
		--descriptor $(DESCRIPTOR) \
		--service liaison-server \
		--port 9080 >$(ENVOYFILTER)

$(ALL_JS_PROTOS):
	$(call protoc_gen_ts,$@,``,$(patsubst %_pb.js,%.proto,$@))

$(ALL_TS_PROTOS):
	$(call protoc_gen_ts,$@,``,$(patsubst %_pb.d.ts,%.proto,$@))

node_modules: touch
	@yarn

.PHONY: js touch
js: node_modules touch $(ALL_JS_PROTOS) $(ALL_TS_PROTOS)

touch:
	mkdir -p ./google/api
	mkdir -p ./validate
	touch ./google/api/annotations_pb.js
	touch ./validate/validate_pb.js

clean:
	rm -fr $(ALL_GENERATED) ./google ./validate
	@echo "cleaned $(ALL_GENERATED)"
