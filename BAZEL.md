# Bazel Notes

The project uses [Bazel](https://bazel.build) as the underlying build system. The [Makefile](Makefile) is provided as a simplified UX over common Bazel commands.

# Go Notes

## Registering a Dependency

Registering a dependency means adding it to the [WORKSPACE](WORKSPACE) file.

Bazel can parse `go.mod` files. The workflow to install a new dependency into the [WORKSPACE](WORKSPACE) is:

1. `go get -u $url` OR edit `go.mod` by hand.
2. `make gazelle-update-repos`

What this does:

1. Downloads the package and puts the proper dependency line into `go.mod`
2. Modifies [WORKSPACE](WORKSPACE) to have a proper `go_repository` entry. 

```python
go_repository(
    name = "org_golang_x_crypto",
    commit = "74369b46fc67",
    importpath = "golang.org/x/crypto",
)
```

## Consuming a Dependency

Main Package: `"@com_github_go_chi_chi//:go_default_library"`
Sub Package: `"@com_github_go_chi_chi//middleware:go_default_library"`

# FAQ

## Why Bazel?

1. I believe in its goal of fast and correct builds. 
2. I also wanted a unified build system for a (potentially) polyglot mono-repo project. I did not want to take on the tech debt of retrofitting later.
3. This project gives me a place to learn Bazel.
