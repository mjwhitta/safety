# safety

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/safety?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/safety)
![License](https://img.shields.io/github/license/mjwhitta/safety?style=for-the-badge)

## What is this?

This Go module provides some thread-safe utilities.

## How to install

Open a terminal and run the following:

```
$ go get -u github.com/mjwhitta/safety
```

## Usage

```
package main

import (
    hl "github.com/mjwhitta/hilighter"
    "github.com/mjwhitta/safety"
)

func main() {
    var m *safety.Map = safety.NewMap()
    var s *safety.Set = safety.NewSet()
    var sl *safety.Slice = safety.NewSlice()

    // Maps
    m.Put("a", "asdf")
    m.Put("b", "blah")
    m.Put("s", "stop")
    m.Put("t", "test")

    m.Range(
        func(k, v any) bool {
            switch v.(string) {
            case "stop":
                return true
            }

            hl.Printf("%s = %s\n", k.(string), v.(string))
            return false
        },
    )

    hl.Println(m.Keys())

    // Sets
    s.Add("asdf")
    s.Add("blah")
    s.Add("stop")
    s.Add("test")

    s.Range(
        func(entry any) bool {
            switch entry.(string) {
            case "stop":
                return true
            }

            hl.Printf("Set includes: %s\n", entry.(string))
            return false
        },
    )

    hl.Println(s.Get())

    // Slices
    sl.Append("asdf")
    sl.Append("blah")
    sl.Append("stop")
    sl.Append("test")

    sl.Range(
        func(i int, v any) bool {
            switch v.(string) {
            case "stop":
                return true
            }

            hl.Printf("Slice %d: %s\n", i, v.(string))
            return false
        },
    )

    hl.Println(sl.Get(0))
}
```

## Links

- [Source](https://github.com/mjwhitta/safety)
