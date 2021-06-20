# go-env-default

[![Go Reference](https://pkg.go.dev/badge/github.com/caitlinelfring/go-env-default.svg)](https://pkg.go.dev/github.com/caitlinelfring/go-env-default)
[![Go Report Card](https://goreportcard.com/badge/github.com/caitlinelfring/go-env-default)](https://goreportcard.com/report/github.com/caitlinelfring/go-env-default)
[![codecov](https://codecov.io/gh/caitlinelfring/go-env-default/branch/main/graph/badge.svg?token=Og3fM2U2HE)](https://codecov.io/gh/caitlinelfring/go-env-default)
[![Go Tests](https://github.com/caitlinelfring/go-env-default/actions/workflows/tests.yaml/badge.svg)](https://github.com/caitlinelfring/go-env-default/actions/workflows/tests.yaml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/caitlinelfring/go-env-default)

```go
import "github.com/caitlinelfring/go-env-default"
```

A collection of helper go functions for accessing the value of
environment variables with pre-defined default values, including type conversion.

Supports the following types:

* `string`
* `bool`
* `int`
* `int64`
* `float64`
* `time.Duration`

Supported Go versions:

* 1.15
* 1.16

## Usage

Example:

```go
package main

import (
  "fmt"

  env "github.com/caitlinelfring/go-env-default"
)

func main() {
  fmt.Println(env.GetDefault("MY_ENV_VAR", "foo"))
}
```

```bash
$ go run main.go
foo
$ MY_ENV_VAR=bar go run main.go
bar
```
