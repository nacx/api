workspace(name = "io_tetrate_api")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Pulling envoy_api repo here to provide basic dependency mapping and common proto dependencies like PGV.
http_archive(
    name = "envoy_api",
    sha256 = "2086d2f8dc6b049005c1efec3ee14cb6bc9328c35436b013801b67e20af40574",
    strip_prefix = "data-plane-api-9b99c017bd625441c2a2f4f682a9cdb03f7c933e",
    urls = [
        "https://github.com/envoyproxy/data-plane-api/archive/9b99c017bd625441c2a2f4f682a9cdb03f7c933e.tar.gz",
    ],
)

load("@envoy_api//bazel:repositories.bzl", envoy_api_dependencies = "api_dependencies")

envoy_api_dependencies()

load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    go = True,
)

load("@com_envoyproxy_protoc_gen_validate//bazel:repositories.bzl", "pgv_dependencies")

pgv_dependencies()

load("@com_envoyproxy_protoc_gen_validate//bazel:dependency_imports.bzl", "pgv_dependency_imports")

pgv_dependency_imports()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "142dd33e38b563605f0d20e89d9ef9eda0fc3cb539a14be1bdb1350de2eda659",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.22.2/rules_go-v0.22.2.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.22.2/rules_go-v0.22.2.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()
