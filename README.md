# safety

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/mjwhitta/safety)](https://goreportcard.com/report/gitlab.com/mjwhitta/safety)

## What is this?

This Go module provides some thread-safe utilities.

## How to install

Open a terminal and run the following:

```
$ go get -u gitlab.com/mjwhitta/safety
```

## Usage

```
package main

import (
    hl "gitlab.com/mjwhitta/hilighter"
    "gitlab.com/mjwhitta/safety"
)

func main() {
    var m *safety.Map = safety.NewMap()

    m.Put("a", "asdf")
    m.Put("b", "blah")
    m.Put("s", "stop")
    m.Put("t", "test")

    m.Range(
        func(k, v interface{}) bool {
            switch v.(string) {
            case "stop":
                return true
            }

            hl.Printf("%s = %s\n", k.(string), v.(string))
            return false
        },
    )

    hl.Println(m.Keys())
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/safety)
