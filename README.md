# 🧰 optkit

A lightweight toolkit for composable Go configuration using typed functional options.

## ✨ Features

- **Type-safe option fields** via generics  
- **Composable option functions** for configuration chaining  
- **Automatic field binding** using reflection + struct tags  
- **Flexible defaults and overrides** via `Options` maps

## 🚀 Getting Started

### Define your options struct

```go
package mypkg

import (
	"github.com/moshenahmias/optkit"
)

type MyOptions struct {
    Foo optkit.Field[int, MyOptions] `id:"foo"`     // without the tag the default is the field name ("Foo")
    Bar optkit.Field[string, MyOptions] `id:"bar"`  // without the tag the default is the field name ("Bar")
}
```

### Initialize it:

```go
package mypkg

import (
	"github.com/moshenahmias/optkit"
)

func init() {
    options := optkit.Init[MyOptions]()
}

func Options() MyOptions {
    return options
}
```

### Write some code that uses it:

```go
package mypkg

import (
	"github.com/moshenahmias/optkit"
)

func Function(options ...optkit.Option[MyOptions]) MyOptions {
    opts := optkit.Build(options...)

    foo := optkit.Get[int](opts, "foo")  // default: 0
    bar := optkit.GetWithDefault(opts, "bar", "!!!") // default: "!!!"

    // with defaults:
    opts := optkit.BuildWithDefaults(optkit.Options{foo: 10}, options...)

    foo := optkit.Get[int](opts, "foo")  // default: 10
    bar := optkit.Get[string](opts, "bar", "") // default: ""

    print(foo)
    print(bar)
}

func Function2(options ...optkit.Option[MyOptions]) MyOptions {
    opts := optkit.BuildWithDefaults(optkit.Options{foo: 10}, options...)

    foo := optkit.Get[int](opts, "foo")  // default: 10
    bar := optkit.Get[string](opts, "bar") // default: ""

    print(foo)
    print(bar)
}
```

### Usage of your package:

```go
package main

import (
    "mypkg"
)

func main() {
    mypkg.Function(mypkg.Options().Foo.Set(42), mypkg.Options().Bar.Set("Hello"))
    // 42
    // Hello

    mypkg.Function2(mypkg.Options().Bar.Set("Hello"), mypkg.Options().Foo.Replace(func(v int) int {
		return v + 10
	}))

    // 20 (default + 10)
    // Hello
}
```
