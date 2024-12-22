# babycli

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/babycli.svg)](https://pkg.go.dev/cattlecloud.net/go/babycli)
[![License](https://img.shields.io/github/license/cattlecloud/babycli?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/babycli/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/babycli/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/babycli/actions/workflows/ci.yaml)

`babycli` is a modern CLI arguments parser for Go.

### Requirements

The minimum Go version is `go1.23`.

### Getting Started

The `babycli` package can be added to a project with `go get`.

```shell
go get cattlecloud.net/go/babycli@latest
```

```go
import "cattlecloud.net/go/babycli"
```

### Influence

This library was made after many projects of using [urfave/cli](https://github.com/urfave/cli),
one of the better command line parser libraries for Go. `babycli` models itself
after this library but with a cleaner, more robust implementation under the hood.
Credit of inspiration belongs to the authors.

### Experimental

(!) Please note this library is still experimental and being worked on, with
breaking changes being very likely.

### License

The `cattlecloud.net/go/babycli` module is open source under the [BSD](LICENSE) license.
