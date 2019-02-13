load("@io_bazel_rules_docker//container:container.bzl", "container_image")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "types.go"
    ],
    importpath = "github.com/plombardi89/http01",
    visibility = ["//visibility:public"],
    deps = [
        #"//pkg/cmd:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
        "@com_github_go_chi_chi//:go_default_library",
        "@com_github_go_chi_chi//middleware:go_default_library",
    ],
)

go_binary(
    name = "tpolis-http01-server",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
