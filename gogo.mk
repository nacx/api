# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

EMPTY :=
SPACE := $(EMPTY) $(EMPTY)

IMPORTMAPS := \
	gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto, \
	google/protobuf/any.proto=github.com/gogo/protobuf/types, \
	google/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor, \
	google/protobuf/duration.proto=github.com/gogo/protobuf/types, \
	google/protobuf/struct.proto=github.com/gogo/protobuf/types, \
	google/protobuf/timestamp.proto=github.com/gogo/protobuf/types, \
	google/protobuf/wrappers.proto=github.com/gogo/protobuf/types, \
	google/rpc/status.proto=istio.io/gogo-genproto/googleapis/google/rpc, \
	google/rpc/code.proto=istio.io/gogo-genproto/googleapis/google/rpc, \
	google/rpc/error_details.proto=istio.io/gogo-genproto/googleapis/google/rpc, \
	google/api/annotations.proto=istio.io/gogo-genproto/googleapis/google/api, \
	google/api/http.proto=istio.io/gogo-genproto/googleapis/google/api, \
	google/protobuf/empty.proto=github.com/gogo/protobuf/types, \
	opencensus/proto/trace/v1/trace.proto=istio.io/gogo-genproto/opencensus/proto/trace/v1, \
	opencensus/proto/trace/v1/trace_config.proto=istio.io/gogo-genproto/opencensus/proto/trace/v1, \
	metrics.proto=istio.io/gogo-genproto/prometheus, \
	validate/validate.proto=github.com/envoyproxy/protoc-gen-validate/validate

MAPS_WITH_SPACE := $(foreach MAP,$(IMPORTMAPS),M$(MAP))
MAPS := $(subst $(SPACE),$(EMPTY),$(MAPS_WITH_SPACE))
