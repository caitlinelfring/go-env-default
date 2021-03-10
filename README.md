# go-env-default

<!-- [![Coverage Status](https://coveralls.io/repos/caitlinelfring/go-env-default/badge.svg?branch=develop&service=github)](https://coveralls.io/github/caitlinelfring/go-env-default?branch=master) -->
[![GoDoc](https://godoc.org/github.com/caitlinelfring/go-env-default?status.svg)](https://godoc.org/github.com/caitlinelfring/go-env-default)
[![Go Report Card](https://goreportcard.com/badge/github.com/caitlinelfring/go-env-default)](https://goreportcard.com/report/github.com/caitlinelfring/go-env-default)

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

## Usage

Example:

```go
import (
  "fmt"

  "github.com/caitlinelfring/go-env-default"
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
