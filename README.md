# Dtype

![Golang 1.15](https://img.shields.io/badge/golang-1.15-blue)
[![Linter](https://github.com/aofdev/dtype/workflows/Linter/badge.svg)](https://github.com/aofdev/dtype/actions)

## ⚙️ Installation

```sh
go get -u github.com/aofdev/dtype
```

## ⚡️ Quickstart

```go
package main

import (
	"fmt"

	d "github.com/aofdev/dtype"
)

const text = ""

func main() {
	value, isNull := d.DefaultStringWithNullable(text)
	fmt.Printf("value: %v isNull: %v", value, isNull)
}

//output: value:  isNull: <nil>
```