load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@envoy_api//bazel:external_proto_deps.bzl",
    "EXTERNAL_PROTO_GO_BAZEL_DEP_MAP",
)
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

_GO_PROTO_SUFFIX = "_go_proto"
_GO_IMPORTPATH_PREFIX = "github.com/tetrateio/api-go"

_COMMON_PROTO_DEPS = [
    "@com_google_protobuf//:any_proto",
    "@com_google_protobuf//:descriptor_proto",
    "@com_google_protobuf//:duration_proto",
    "@com_google_protobuf//:empty_proto",
    "@com_google_protobuf//:struct_proto",
    "@com_google_protobuf//:timestamp_proto",
    "@com_google_protobuf//:wrappers_proto",
    "@com_google_googleapis//google/api:http_proto",
    "@com_google_googleapis//google/api:httpbody_proto",
    "@com_google_googleapis//google/api:annotations_proto",
    "@com_google_googleapis//google/rpc:status_proto",
    "@com_envoyproxy_protoc_gen_validate//validate:validate_proto",
]

def _proto_mapping(dep, proto_dep_map, proto_suffix):
    mapped = proto_dep_map.get(dep)
    if mapped == None:
        prefix = "@" + Label(dep).workspace_name if not dep.startswith("//") else ""
        return prefix + "//" + Label(dep).package + ":" + Label(dep).name + proto_suffix
    return mapped

def _go_proto_mapping(dep):
    return _proto_mapping(dep, EXTERNAL_PROTO_GO_BAZEL_DEP_MAP, _GO_PROTO_SUFFIX)

# API proto package
def api_proto_package(
        name = "pkg",
        srcs = [],
        deps = [],
        has_services = False,
        visibility = ["//visibility:public"]):
    """Defines a API proto package.

    Args:
      name: package name
      srcs: source files if not specified all .proto files in the directory
      deps: dependencies
      has_services: whether .proto file has services
      visibility: Bazel visibility
    """
    if srcs == []:
        srcs = native.glob(["*.proto"])

    proto_library(
        name = name,
        srcs = srcs,
        deps = deps + _COMMON_PROTO_DEPS,
        visibility = visibility,
    )

    compilers = ["@io_bazel_rules_go//proto:go_proto", "@envoy_api//bazel:pgv_plugin_go"]
    if has_services:
        compilers = ["@io_bazel_rules_go//proto:go_grpc", "@envoy_api//bazel:pgv_plugin_go"]
    go_proto_library(
        name = name + _GO_PROTO_SUFFIX,
        compilers = compilers,
        importpath = _GO_IMPORTPATH_PREFIX + native.package_name(),
        proto = name,
        visibility = ["//visibility:public"],
        deps = [_go_proto_mapping(dep) for dep in deps] + [
            "@com_github_golang_protobuf//ptypes:go_default_library",
            "@com_github_golang_protobuf//ptypes/any:go_default_library",
            "@com_github_golang_protobuf//ptypes/duration:go_default_library",
            "@com_github_golang_protobuf//ptypes/struct:go_default_library",
            "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",
            "@com_github_golang_protobuf//ptypes/wrappers:go_default_library",
            "@com_envoyproxy_protoc_gen_validate//validate:go_default_library",
            "@com_google_googleapis//google/api:annotations_go_proto",
            "@com_google_googleapis//google/rpc:status_go_proto",
        ],
    )
