# 🧰 optkit

A lightweight toolkit for composable Go configuration using typed functional options.

## ✨ Features

- **Type-safe option fields** via generics  
- **Composable option functions** for configuration chaining  
- **Automatic field binding** using reflection + struct tags  
- **Flexible defaults and overrides** via `Options` maps

## 🚀 Getting Started

### Define your options struct:

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

var (
    options = optkit.Init[MyOptions]()
)

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

func Function(options ...optkit.Option[MyOptions]) {
    opts := optkit.Build(options...)

    foo := optkit.Get[int](opts, "foo")  // default: 0
    bar := optkit.GetWithDefault(opts, "bar", "!!!") // default: "!!!"

    // with defaults:
    opts := optkit.BuildWithDefaults(optkit.Options{"foo": 10}, options...)

    foo := optkit.Get[int](opts, "foo")  // default: 10
    bar := optkit.Get[string](opts, "bar", "") // default: ""

    print(foo)
    print(bar)
}

func Function2(options ...optkit.Option[MyOptions]) {
    opts := optkit.BuildWithDefaults(optkit.Options{"foo": 10}, options...)

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

    mypkg.Function2(mypkg.Options().Bar.Set("Hello"), mypkg.Options().Foo.Replace(func(v *int) {
        *v = *v + 10
    }))

    // 20 (default + 10)
    // Hello
}
```

### 🧩 Multi-Value Fields (Field2 to Field6)

The optkit package supports type-safe functional options that can hold multiple values under a single field key:

```go
type Config struct {
    Limits optkit.Field2[int, string, Config]
}

config := optkit.Init[Config]()

opts := optkit.Build(
    config.Limits.Set(42, "max"),
)

min, label := optkit.Get2[int, string](opts, "Limits")

fmt.Println(min, label) // Output: 42 max
```

You can also use .Replace() to modify both values in-place:

```go
opts = optkit.Build(
    config.Limits.Set(10, "low"),
    config.Limits.Replace(func(i *int, s *string) {
        *i = *i * 2
        *s = strings.ToUpper(*s)
    }),
)
```

### 🌟 Vars

OptKit also exposes a lower-level typed API using Var, Var2, ..., Var6 types. These give you direct, generic access to the same composable option pattern without needing to bind them to struct fields, but you do need to bind them to some type defined in your package - the second type parameter S is exactly what anchors them to a particular context or package. That type parameter acts as a "scope tag", ensuring type isolation across packages and uses.

```go
package mypkg

type FuncThreeOptions struct{}
type SomeOtherOptions struct{}

const ( 
    Foo optkit.Var[int, FuncThreeOptions] = "foo"
    Bar optkit.Var[string, FuncThreeOptions] = "bar"
    ... optkit.Var2[..., ..., SomeOtherOptions] = "..." // Var2
    ... optkit.Var6[..., ..., ..., ..., ..., ..., SomeOtherOptions] = "..." // Var6
)

func Function3(options ...optkit.Option[FuncThreeOptions]) {
    opts := optkit.Build(options...)

    foo := optkit.Get[int](opts, "foo")  // default: 0
    bar := optkit.GetWithDefault(opts, "bar", "!!!") // default: "!!!"

    print(foo)
    print(bar)
}
```

Then:

```go
func main() {
    mypkg.Function3(mypkg.Foo.Set(42), mypkg.Bar.Set("Hello"))
    // 42
    // Hello
}
```