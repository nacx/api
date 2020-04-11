# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

all: gen

# General setup.
repo_dir := .
out_path = /tmp

protoc = protoc -Icommon-protos/github.com/envoyproxy/protoc-gen-validate -Icommon-protos -I.
prototool = prototool

go_plugin_prefix := --go_out=plugins=grpc,
go_plugin := $(go_plugin_prefix):$(out_path)

# Setup for protoc_gen_*.
gogofast_plugin_prefix := --gogofast_out=plugins=grpc,
gogoslick_plugin_prefix := --gogoslick_out=plugins=grpc,
grpc_gateway_plugin_prefix := --grpc-gateway_out=allow_delete_body=true,
protoc_gen_validate_plugin_prefix := --validate_out=lang=gogo

# Setup for doc.
protoc_gen_docs_plugin := --docs_out=warnings=true,dictionary=$(repo_dir)/dictionaries/en-US,custom_word_list=$(repo_dir)/dictionaries/custom.txt,mode=html_fragment_with_front_matter:$(repo_dir)/
protoc_gen_docs_plugin_per_file := --docs_out=warnings=true,dictionary=$(repo_dir)/dictionaries/en-US,custom_word_list=$(repo_dir)/dictionaries/custom.txt,per_file=true,mode=html_fragment_with_front_matter:$(repo_dir)/

empty :=
space := $(empty) $(empty)

importmaps := \
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

mapping_with_spaces := $(foreach mapping,$(importmaps),M$(mapping))
gogo_mapping := $(subst $(space),$(empty),$(mapping_with_spaces))

gogofast_plugin := $(gogofast_plugin_prefix)$(gogo_mapping):$(out_path)
gogoslick_plugin := $(gogoslick_plugin_prefix)$(gogo_mapping):$(out_path)
grpc_gateway_plugin := $(grpc_gateway_plugin_prefix)$(gogo_mapping):$(out_path)
protoc_gen_validate_plugin := $(protoc_gen_validate_plugin_prefix):$(out_path)

# Generation rules.

gen: \
	generate-audit \
	generate-authz \
	generate-operator-common \
	generate-operator-managementplane \
	generate-operator-controlplane \
	generate-q-rbac \
	generate-regsource \
	generate-settings \
	generate-spm-alert \
	generate-spm-config \
	generate-spm-metadata \
	generate-spm-notification \
	generate-tcc-core \
	generate-tcc-workflows-loadbalancer \
	generate-test \
	generate-troubleshoot \
	generate-tsb-gateway \
	generate-tsb-rbac \
	generate-tsb-security \
	generate-tsb-traffic \
	generate-tsb-service \
	generate-tsb-private \
	generate-tsb-types \
	generate-tsb \
	tidy-go

clean: \
	clean-audit \
	clean-authz \
	clean-operator-common \
	clean-operator-managementplane \
	clean-operator-controlplane \
	clean-q-rbac \
	clean-regsource \
	clean-settings \
	clean-spm-alert \
	clean-spm-config \
	clean-spm-metadata \
	clean-spm-notification \
	clean-tcc-core \
	clean-tcc-workflows-loadbalancer \
	clean-test \
	clean-troubleshoot \
	clean-tsb-gateway \
	clean-tsb-rbac \
	clean-tsb-security \
	clean-tsb-traffic \
	clean-tsb-service \
	clean-tsb-private \
	clean-tsb-types \
	clean-tsb

# audit/...
audit_v1_path := audit/v1
audit_v1_protos := $(wildcard $(audit_v1_path)/*.proto)
audit_v1_pb_gos := $(audit_v1_protos:.proto=.pb.go)
audit_v1_pb_gw_gos := $(audit_v1_protos:.proto=.pb.gw.go)
audit_v1_pb_validate_gos := $(audit_v1_protos:.proto=.pb.validate.go)
audit_v1_pb_doc := $(audit_v1_path)/tetrate.api.audit.v1.pb.html

$(audit_v1_pb_gos) $(audit_v1_pb_gw_gos) $(audit_v1_pb_validate_gos) $(audit_v1_pb_doc): $(audit_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(audit_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/audit/* audit

generate-audit: $(audit_v1_pb_gos)

clean-audit:
	@rm -fr $(audit_v1_pb_gos) $(audit_v1_pb_gw_gos) $(audit_v1_pb_validate_gos) $(audit_v1_pb_doc)

# authz/...
authz_v1_path := authz/v1
authz_v1_protos := $(wildcard $(authz_v1_path)/*.proto)
authz_v1_pb_gos := $(authz_v1_protos:.proto=.pb.go)
authz_v1_pb_gw_gos := $(authz_v1_protos:.proto=.pb.gw.go)
authz_v1_pb_validate_gos := $(authz_v1_protos:.proto=.pb.validate.go)
authz_v1_pb_doc := $(authz_v1_path)/tetrate.api.authz.v1.pb.html

$(authz_v1_pb_gos) $(authz_v1_pb_gw_gos) $(authz_v1_pb_validate_gos) $(authz_v1_pb_doc): $(authz_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(authz_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/authz/* authz

generate-authz: $(authz_v1_pb_gos)

clean-authz:
	@rm -fr $(authz_v1_pb_gos) $(authz_v1_pb_gw_gos) $(authz_v1_pb_validate_gos)

# operator/...
# operator/common/...
operator_common_path := operator/common
operator_common_protos := $(wildcard $(operator_common_path)/*.proto)
operator_common_pb_gos := $(operator_common_protos:.proto=.pb.go)
operator_common_pb_doc := $(operator_common_path)/tetrate.api.operator.common.pb.html

$(operator_common_pb_gos) $(operator_common_pb_doc): $(operator_common_protos)
	@$(protoc) $(go_plugin) $(protoc_gen_docs_plugin)$(operator_common_path) $^
	@cp -r /tmp/github.com/tetrateio/api/operator/common/* operator/common
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(operator_common_path)/kubernetes.pb.go

generate-operator-common: $(operator_common_pb_gos)

clean-operator-common:
	@rm -fr $(operator_common_pb_gos)

# operator/managementplane/...
operator_managementplane_v1alpha1_path := operator/managementplane/v1alpha1
operator_managementplane_v1alpha1_protos := $(wildcard $(operator_managementplane_v1alpha1_path)/*.proto)
operator_managementplane_v1alpha1_pb_gos := $(operator_managementplane_v1alpha1_protos:.proto=.pb.go)
operator_managementplane_v1alpha1_pb_doc := $(operator_managementplane_v1alpha1_path)/tetrate.api.operator.managementplane.v1alpha1.pb.html

$(operator_managementplane_v1alpha1_pb_gos) $(operator_managementplane_v1alpha1_pb_doc): $(operator_managementplane_v1alpha1_protos)
	@$(protoc) $(go_plugin) $(protoc_gen_docs_plugin)$(operator_managementplane_v1alpha1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/operator/managementplane/* operator/managementplane
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(operator_managementplane_v1alpha1_path)/values_types.pb.go
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(operator_managementplane_v1alpha1_path)/component.pb.go
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(operator_managementplane_v1alpha1_path)/operator.pb.go

generate-operator-managementplane: $(operator_managementplane_v1alpha1_pb_gos)

clean-operator-managementplane:
	@rm -fr $(operator_managementplane_v1alpha1_pb_gos)


# operator/controlplane/...
operator_controlplane_v1alpha1_path := operator/controlplane/v1alpha1
operator_controlplane_v1alpha1_protos := $(wildcard $(operator_controlplane_v1alpha1_path)/*.proto)
operator_controlplane_v1alpha1_pb_gos := $(operator_controlplane_v1alpha1_protos:.proto=.pb.go)
operator_controlplane_v1alpha1_pb_doc := $(operator_controlplane_v1alpha1_path)/tetrate.api.operator.controlplane.v1alpha1.pb.html

$(operator_controlplane_v1alpha1_pb_gos) $(operator_controlplane_v1alpha1_pb_doc): $(operator_controlplane_v1alpha1_protos)
	@$(protoc) $(go_plugin) $(protoc_gen_docs_plugin)$(operator_controlplane_v1alpha1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/operator/controlplane/* operator/controlplane
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(operator_controlplane_v1alpha1_path)/controlplane_component.pb.go
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(operator_controlplane_v1alpha1_path)/controlplane.pb.go

generate-operator-controlplane: $(operator_controlplane_v1alpha1_pb_gos)

clean-operator-controlplane:
	@rm -fr $(operator_controlplane_v1alpha1_pb_gos)

# q/...
# q/rbac/...
q_rbac_v1_path := q/rbac/v1
q_rbac_v1_protos := $(wildcard $(q_rbac_v1_path)/*.proto)
q_rbac_v1_pb_gos := $(q_rbac_v1_protos:.proto=.pb.go)
q_rbac_v1_pb_gw_gos := $(q_rbac_v1_protos:.proto=.pb.gw.go)
q_rbac_v1_pb_validate_gos := $(q_rbac_v1_protos:.proto=.pb.validate.go)
q_rbac_v1_pb_doc := $(q_rbac_v1_path)/tetrate.api.q.rbac.v1.pb.html

$(q_rbac_v1_pb_gos) $(q_rbac_v1_pb_gw_gos) $(q_rbac_v1_pb_validate_gos) $(q_rbac_v1_pb_doc): $(q_rbac_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(q_rbac_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/q/rbac/* q/rbac

generate-q-rbac: $(q_rbac_v1_pb_gos)

clean-q-rbac:
	@rm -fr $(q_rbac_v1_pb_gos) $(q_rbac_v1_pb_gw_gos) $(q_rbac_v1_pb_validate_gos)

# regsource/...
regsource_v1_path := regsource/v1
regsource_v1_protos := $(wildcard $(regsource_v1_path)/*.proto)
regsource_v1_pb_gos := $(regsource_v1_protos:.proto=.pb.go)
regsource_v1_pb_gw_gos := $(regsource_v1_protos:.proto=.pb.gw.go)
regsource_v1_pb_validate_gos := $(regsource_v1_protos:.proto=.pb.validate.go)
regsource_v1_pb_doc := $(regsource_v1_path)/tetrate.api.regsource.v1.pb.html

$(regsource_v1_pb_gos) $(regsource_v1_pb_gw_gos) $(regsource_v1_pb_validate_gos) $(regsource_v1_pb_doc): $(regsource_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(regsource_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/regsource/* regsource

generate-regsource: $(regsource_v1_pb_gos)

clean-regsource:
	@rm -fr $(regsource_v1_pb_gos) $(regsource_v1_pb_gw_gos) $(regsource_v1_pb_validate_gos)

# settings/...
settings_v1_path := settings/v1
settings_v1_protos := $(wildcard $(settings_v1_path)/*.proto)
settings_v1_pb_gos := $(settings_v1_protos:.proto=.pb.go)
settings_v1_pb_gw_gos := $(settings_v1_protos:.proto=.pb.gw.go)
settings_v1_pb_validate_gos := $(settings_v1_protos:.proto=.pb.validate.go)
settings_v1_pb_doc := $(settings_v1_path)/tetrate.api.settings.v1.pb.html

$(settings_v1_pb_gos) $(settings_v1_pb_gw_gos) $(settings_v1_pb_validate_gos) $(settings_v1_pb_doc): $(settings_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(settings_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/settings/* settings

generate-settings: $(settings_v1_pb_gos) $(settings_v1_pb_gw_gos) $(settings_v1_pb_validate_gos)

clean-settings:
	@rm -fr $(settings_v1_pb_gos) $(settings_v1_pb_gw_gos) $(settings_v1_pb_validate_gos)

# spm/alert/...
spm_alert_v1_path := spm/alert/v1
spm_alert_v1_protos := $(wildcard $(spm_alert_v1_path)/*.proto)
spm_alert_v1_pb_gos := $(spm_alert_v1_protos:.proto=.pb.go)
spm_alert_v1_pb_gw_gos := $(spm_alert_v1_protos:.proto=.pb.gw.go)
spm_alert_v1_pb_validate_gos := $(spm_alert_v1_protos:.proto=.pb.validate.go)
spm_alert_v1_pb_doc := $(spm_alert_v1_path)/tetrate.api.spm.alert.v1.pb.html

$(spm_alert_v1_pb_gos) $(spm_alert_v1_pb_gw_gos) $(spm_alert_v1_pb_validate_gos) $(spm_alert_v1_pb_doc): $(spm_alert_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(spm_alert_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/spm/alert/* spm/alert

generate-spm-alert: $(spm_alert_v1_pb_gos)

clean-spm-alert:
	@rm -fr $(spm_alert_v1_pb_gos) $(spm_alert_v1_pb_gw_gos) $(spm_alert_v1_pb_validate_gos) $(spm_alert_v1_pb_doc)

# spm/config/...
spm_config_v1_path := spm/config/v1
spm_config_v1_protos := $(wildcard $(spm_config_v1_path)/*.proto)
spm_config_v1_pb_gos := $(spm_config_v1_protos:.proto=.pb.go)
spm_config_v1_pb_gw_gos := $(spm_config_v1_protos:.proto=.pb.gw.go)
spm_config_v1_pb_validate_gos := $(spm_config_v1_protos:.proto=.pb.validate.go)
spm_config_v1_pb_doc := $(spm_config_v1_path)/tetrate.api.spm.config.v1.pb.html

$(spm_config_v1_pb_gos) $(spm_config_v1_pb_gw_gos) $(spm_config_v1_pb_validate_gos) $(spm_config_v1_pb_doc): $(spm_config_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(spm_config_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/spm/config/* spm/config

generate-spm-config: $(spm_config_v1_pb_gos)

clean-spm-config:
	@rm -fr $(spm_config_v1_pb_gos) $(spm_config_v1_pb_gw_gos) $(spm_config_v1_pb_validate_gos) $(spm_config_v1_pb_doc)

# spm/metadata/...
spm_metadata_v1_path := spm/metadata/v1
spm_metadata_v1_protos := $(wildcard $(spm_metadata_v1_path)/*.proto)
spm_metadata_v1_pb_gos := $(spm_metadata_v1_protos:.proto=.pb.go)
spm_metadata_v1_pb_gw_gos := $(spm_metadata_v1_protos:.proto=.pb.gw.go)
spm_metadata_v1_pb_validate_gos := $(spm_metadata_v1_protos:.proto=.pb.validate.go)
spm_metadata_v1_pb_doc := $(spm_metadata_v1_path)/tetrate.api.spm.metadata.v1.pb.html

$(spm_metadata_v1_pb_gos) $(spm_metadata_v1_pb_gw_gos) $(spm_metadata_v1_pb_validate_gos) $(spm_metadata_v1_pb_doc): $(spm_metadata_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(spm_metadata_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/spm/metadata/* spm/metadata

generate-spm-metadata: $(spm_metadata_v1_pb_gos)

clean-spm-metadata:
	@rm -fr $(spm_metadata_v1_pb_gos) $(spm_metadata_v1_pb_gw_gos) $(spm_metadata_v1_pb_validate_gos) $(spm_metadata_v1_pb_doc)

# spm/notification/...
spm_notification_v1_path := spm/notification/v1
spm_notification_v1_protos := $(wildcard $(spm_notification_v1_path)/*.proto)
spm_notification_v1_pb_gos := $(spm_notification_v1_protos:.proto=.pb.go)
spm_notification_v1_pb_gw_gos := $(spm_notification_v1_protos:.proto=.pb.gw.go)
spm_notification_v1_pb_validate_gos := $(spm_notification_v1_protos:.proto=.pb.validate.go)
spm_notification_v1_pb_doc := $(spm_notification_v1_path)/tetrate.api.spm.notification.v1.pb.html

$(spm_notification_v1_pb_gos) $(spm_notification_v1_pb_gw_gos) $(spm_notification_v1_pb_validate_gos) $(spm_notification_v1_pb_doc): $(spm_notification_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(spm_notification_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/spm/notification/* spm/notification

generate-spm-notification: $(spm_notification_v1_pb_gos)

clean-spm-notification:
	@rm -fr $(spm_notification_v1_pb_gos) $(spm_notification_v1_pb_gw_gos) $(spm_notification_v1_pb_validate_gos) $(spm_notification_v1_pb_doc)

# tcc/...
# tcc/core/...
tcc_core_v1_path := tcc/core/v1
tcc_core_v1_protos := $(wildcard $(tcc_core_v1_path)/*.proto)
tcc_core_v1_pb_gos := $(tcc_core_v1_protos:.proto=.pb.go)
tcc_core_v1_pb_gw_gos := $(tcc_core_v1_protos:.proto=.pb.gw.go)
tcc_core_v1_pb_validate_gos := $(tcc_core_v1_protos:.proto=.pb.validate.go)
tcc_core_v1_pb_doc := $(tcc_core_v1_path)/tetrate.api.tcc.core.v1.pb.html

$(tcc_core_v1_pb_gos) $(tcc_core_v1_pb_gw_gos) $(tcc_core_v1_pb_validate_gos) $(tcc_core_v1_pb_doc): $(tcc_core_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(tcc_core_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tcc/core/* tcc/core

generate-tcc-core: $(tcc_core_v1_pb_gos)

clean-tcc-core:
	@rm -fr $(tcc_core_v1_pb_gos) $(tcc_core_v1_pb_gw_gos) $(tcc_core_v1_pb_validate_gos) $(tcc_core_v1_pb_doc)

# tcc/workflows/...
# tcc/workflows/loadbalancer/...
tcc_workflows_loadbalancer_v1_path := tcc/workflows/loadbalancer/v1
tcc_workflows_loadbalancer_v1_protos := $(wildcard $(tcc_workflows_loadbalancer_v1_path)/*.proto)
tcc_workflows_loadbalancer_v1_pb_gos := $(tcc_workflows_loadbalancer_v1_protos:.proto=.pb.go)
tcc_workflows_loadbalancer_v1_pb_gw_gos := $(tcc_workflows_loadbalancer_v1_protos:.proto=.pb.gw.go)
tcc_workflows_loadbalancer_v1_pb_validate_gos := $(tcc_workflows_loadbalancer_v1_protos:.proto=.pb.validate.go)
tcc_workflows_loadbalancer_v1_doc := $(tcc_workflows_loadbalancer_v1_path)/tetrate.api.tcc.workflows.loadbalancer.v1.pb.html

$(tcc_workflows_loadbalancer_v1_pb_gos) $(tcc_workflows_loadbalancer_v1_pb_gw_gos) $(tcc_workflows_loadbalancer_v1_pb_validate_gos) $(tcc_workflows_loadbalancer_v1_doc): $(tcc_workflows_loadbalancer_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(tcc_workflows_loadbalancer_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tcc/workflows/loadbalancer/* tcc/workflows/loadbalancer

generate-tcc-workflows-loadbalancer: $(tcc_workflows_loadbalancer_v1_pb_gos)

clean-tcc-workflows-loadbalancer:
	@rm -fr $(tcc_workflows_loadbalancer_v1_pb_gos) $(tcc_workflows_loadbalancer_v1_pb_gw_gos) $(tcc_workflows_loadbalancer_v1_pb_validate_gos) $(tcc_workflows_loadbalancer_v1_doc)

# test/...
test_path := test/v1
test_protos := $(wildcard $(test_path)/*.proto)
test_pb_gos := $(test_protos:.proto=.pb.go)
test_pb_gw_gos := $(test_protos:.proto=.pb.gw.go)
test_pb_validate_gos := $(test_protos:.proto=.pb.validate.go)

$(test_pb_gos) $(test_pb_gw_gos) $(test_pb_validate_gos): $(test_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $^
	@cp -r /tmp/github.com/tetrateio/api/test/* test

generate-test: $(test_pb_gos)

clean-test:
	@rm -fr $(test_pb_gos) $(test_pb_gw_gos) $(test_pb_validate_gos)

# troubleshoot/...
troubleshoot_v1_path := troubleshoot/v1
troubleshoot_v1_protos := $(wildcard $(troubleshoot_v1_path)/*.proto)
troubleshoot_v1_pb_gos := $(troubleshoot_v1_protos:.proto=.pb.go)
troubleshoot_v1_pb_gw_gos := $(troubleshoot_v1_protos:.proto=.pb.gw.go)
troubleshoot_v1_pb_validate_gos := $(troubleshoot_v1_protos:.proto=.pb.validate.go)
troubleshoot_v1_pb_doc := $(troubleshoot_v1_path)/tetrate.api.troubleshoot.v1.pb.html

$(troubleshoot_v1_pb_gos) $(troubleshoot_v1_pb_gw_gos) $(troubleshoot_v1_pb_validate_gos) $(troubleshoot_v1_pb_doc): $(troubleshoot_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin)$(troubleshoot_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/troubleshoot/* troubleshoot

generate-troubleshoot: $(troubleshoot_v1_pb_gos)

clean-troubleshoot:
	@rm -fr $(troubleshoot_v1_pb_gos) $(troubleshoot_v1_pb_gw_gos) $(troubleshoot_v1_pb_validate_gos) $(troubleshoot_v1_pb_doc)

# tsb/...
tsb_v1_path := tsb/v1
tsb_v1_protos := $(wildcard $(tsb_v1_path)/*.proto)
tsb_v1_pb_gos := $(tsb_v1_protos:.proto=.pb.go)
tsb_v1_pb_gw_gos := $(tsb_v1_protos:.proto=.pb.gw.go)
tsb_v1_pb_validate_gos := $(tsb_v1_protos:.proto=.pb.validate.go)
tsb_v1_pb_doc := $(tsb_v1_path)/tetrate.api.tsb.v1.pb.html

$(tsb_v1_pb_gos) $(tsb_v1_pb_gw_gos) $(tsb_v1_pb_validate_gos) $(tsb_v1_pb_doc): $(tsb_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin_per_file)$(tsb_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tsb/* tsb

generate-tsb: $(tsb_v1_pb_gos)

clean-tsb:
	@rm -fr $(tsb_v1_pb_gos) $(tsb_v1_pb_gw_gos) $(tsb_v1_pb_validate_gos) $(tsb_v1_pb_doc)

# tsb/gateway/...
tsb_gateway_v1_path := tsb/gateway/v1
tsb_gateway_v1_protos := $(wildcard $(tsb_gateway_v1_path)/*.proto)
tsb_gateway_v1_pb_gos := $(tsb_gateway_v1_protos:.proto=.pb.go)
tsb_gateway_v1_pb_gw_gos := $(tsb_gateway_v1_protos:.proto=.pb.gw.go)
tsb_gateway_v1_pb_validate_gos := $(tsb_gateway_v1_protos:.proto=.pb.validate.go)
tsb_gateway_v1_pb_doc := $(tsb_gateway_v1_path)/tetrate.api.tsb.gateway.v1.pb.html

$(tsb_gateway_v1_pb_gos) $(tsb_gateway_v1_pb_gw_gos) $(tsb_gateway_v1_pb_validate_gos) $(tsb_gateway_v1_pb_doc): $(tsb_gateway_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin_per_file)$(tsb_gateway_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tsb/gateway/* tsb/gateway

generate-tsb-gateway: $(tsb_gateway_v1_pb_gos)

clean-tsb-gateway:
	@rm -fr $(tsb_gateway_v1_pb_gos) $(tsb_gateway_v1_pb_gw_gos) $(tsb_gateway_v1_pb_validate_gos) $(tsb_gateway_v1_pb_doc)

# tsb/rbac/...
tsb_rbac_v1_path := tsb/rbac/v1
tsb_rbac_v1_protos := $(wildcard $(tsb_rbac_v1_path)/*.proto)
tsb_rbac_v1_pb_gos := $(tsb_rbac_v1_protos:.proto=.pb.go)
tsb_rbac_v1_pb_gw_gos := $(tsb_rbac_v1_protos:.proto=.pb.gw.go)
tsb_rbac_v1_pb_validate_gos := $(tsb_rbac_v1_protos:.proto=.pb.validate.go)
tsb_rbac_v1_pb_doc := $(tsb_rbac_v1_path)/tetrate.api.tsb.gateway.v1.pb.html

$(tsb_rbac_v1_pb_gos) $(tsb_rbac_v1_pb_gw_gos) $(tsb_rbac_v1_pb_validate_gos) $(tsb_rbac_v1_pb_doc): $(tsb_rbac_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin_per_file)$(tsb_rbac_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tsb/rbac/* tsb/rbac

generate-tsb-rbac: $(tsb_rbac_v1_pb_gos)

clean-tsb-rbac:
	@rm -fr $(tsb_rbac_v1_pb_gos) $(tsb_rbac_v1_pb_gw_gos) $(tsb_rbac_v1_pb_validate_gos) $(tsb_rbac_v1_pb_doc)

# tsb/security/...
tsb_security_v1_path := tsb/security/v1
tsb_security_v1_protos := $(wildcard $(tsb_security_v1_path)/*.proto)
tsb_security_v1_pb_gos := $(tsb_security_v1_protos:.proto=.pb.go)
tsb_security_v1_pb_gw_gos := $(tsb_security_v1_protos:.proto=.pb.gw.go)
tsb_security_v1_pb_validate_gos := $(tsb_security_v1_protos:.proto=.pb.validate.go)
tsb_security_v1_pb_doc := $(tsb_security_v1_path)/tetrate.api.tsb.security.v1.pb.html

$(tsb_security_v1_pb_gos) $(tsb_security_v1_pb_gw_gos) $(tsb_security_v1_pb_validate_gos) $(tsb_security_v1_pb_doc): $(tsb_security_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin_per_file)$(tsb_security_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tsb/security/* tsb/security

generate-tsb-security: $(tsb_security_v1_pb_gos)

clean-tsb-security:
	@rm -fr $(tsb_security_v1_pb_gos) $(tsb_security_v1_pb_gw_gos) $(tsb_security_v1_pb_validate_gos) $(tsb_security_v1_pb_doc)

# tsb/traffic/...
tsb_traffic_v1_path := tsb/traffic/v1
tsb_traffic_v1_protos := $(wildcard $(tsb_traffic_v1_path)/*.proto)
tsb_traffic_v1_pb_gos := $(tsb_traffic_v1_protos:.proto=.pb.go)
tsb_traffic_v1_pb_gw_gos := $(tsb_traffic_v1_protos:.proto=.pb.gw.go)
tsb_traffic_v1_pb_validate_gos := $(tsb_traffic_v1_protos:.proto=.pb.validate.go)
tsb_traffic_v1_pb_doc := $(tsb_traffic_v1_path)/tetrate.api.tsb.traffic.v1.pb.html

$(tsb_traffic_v1_pb_gos) $(tsb_traffic_v1_pb_gw_gos) $(tsb_traffic_v1_pb_validate_gos) $(tsb_traffic_v1_pb_doc): $(tsb_traffic_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin_per_file)$(tsb_traffic_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tsb/traffic/* tsb/traffic

generate-tsb-traffic: $(tsb_traffic_v1_pb_gos)

clean-tsb-traffic:
	@rm -fr $(tsb_traffic_v1_pb_gos) $(tsb_traffic_v1_pb_gw_gos) $(tsb_traffic_v1_pb_validate_gos) $(tsb_traffic_v1_pb_doc)

# tsb/types/...
tsb_types_v1_path := tsb/types/v1
tsb_types_v1_protos := $(wildcard $(tsb_types_v1_path)/*.proto)
tsb_types_v1_pb_gos := $(tsb_types_v1_protos:.proto=.pb.go)
tsb_types_v1_pb_gw_gos := $(tsb_types_v1_protos:.proto=.pb.gw.go)
tsb_types_v1_pb_validate_gos := $(tsb_types_v1_protos:.proto=.pb.validate.go)
tsb_types_v1_pb_doc := $(tsb_types_v1_path)/tetrate.api.tsb.types.v1.pb.html

$(tsb_types_v1_pb_gos) $(tsb_types_v1_pb_gw_gos) $(tsb_types_v1_pb_validate_gos) $(tsb_types_v1_pb_doc): $(tsb_types_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin_per_file)$(tsb_types_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tsb/types/* tsb/types

generate-tsb-types: $(tsb_types_v1_pb_gos)

clean-tsb-types:
	@rm -fr $(tsb_types_v1_pb_gos) $(tsb_types_v1_pb_gw_gos) $(tsb_types_v1_pb_validate_gos) $(tsb_types_v1_pb_doc)

# tsb/service/...
tsb_service_v1_path := tsb/service/v1
tsb_service_v1_protos := $(wildcard $(tsb_service_v1_path)/*.proto)
tsb_service_v1_pb_gos := $(tsb_service_v1_protos:.proto=.pb.go)
tsb_service_v1_pb_gw_gos := $(tsb_service_v1_protos:.proto=.pb.gw.go)
tsb_service_v1_pb_validate_gos := $(tsb_service_v1_protos:.proto=.pb.validate.go)
tsb_service_v1_pb_doc := $(tsb_service_v1_path)/tetrate.api.tsb.service.v1.pb.html

$(tsb_service_v1_pb_gos) $(tsb_service_v1_pb_gw_gos) $(tsb_service_v1_pb_validate_gos) $(tsb_service_v1_pb_doc): $(tsb_service_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $(protoc_gen_docs_plugin_per_file)$(tsb_service_v1_path) $^
	@cp -r /tmp/github.com/tetrateio/api/tsb/service/* tsb/service

generate-tsb-service: $(tsb_service_v1_pb_gos)

clean-tsb-service:
	@rm -fr $(tsb_service_v1_pb_gos) $(tsb_service_v1_pb_gw_gos) $(tsb_service_v1_pb_validate_gos) $(tsb_service_v1_pb_doc)

# tsb/private/...
tsb_private_v1_path := tsb/private/v1
tsb_private_v1_protos := $(wildcard $(tsb_private_v1_path)/*.proto)
tsb_private_v1_pb_gos := $(tsb_private_v1_protos:.proto=.pb.go)
tsb_private_v1_pb_gw_gos := $(tsb_private_v1_protos:.proto=.pb.gw.go)
tsb_private_v1_pb_validate_gos := $(tsb_private_v1_protos:.proto=.pb.validate.go)
# No html docs for private protos. We dont want this in the web.

$(tsb_private_v1_pb_gos) $(tsb_private_v1_pb_gw_gos) $(tsb_private_v1_pb_validate_gos): $(tsb_private_v1_protos)
	@$(protoc) $(gogofast_plugin) $(grpc_gateway_plugin) $(protoc_gen_validate_plugin) $^
	@cp -r /tmp/github.com/tetrateio/api/tsb/private/* tsb/private

generate-tsb-private: $(tsb_private_v1_pb_gos)

clean-tsb-private:
	@rm -fr $(tsb_private_v1_pb_gos) $(tsb_private_v1_pb_gw_gos) $(tsb_private_v1_pb_validate_gos)

# Include common targets.
include common/Makefile.common.mk
