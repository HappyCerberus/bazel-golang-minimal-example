load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "hello",
    srcs = ["hello.go"],
    deps = [":greeter"],
)

go_library(
    name = "greeter",
    importpath = "github.com/HappyCerberus/bazel-golang-minimal-example/greeter",
    srcs = ["greeter.go"],
)

go_test(
    name = "greeter_test",
    srcs = [ "greeter_test.go" ],
    embed = [ ":greeter" ],
)
