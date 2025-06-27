package optkit

import (
	"maps"
	"reflect"
	"unsafe"
)

// Init initializes an option set struct.
func Init[T any]() T {
	set := new(T)

	v := reflect.ValueOf(set).Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		panic("optkit: Init requires a pointer to a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.CanSet() && field.Kind() == reflect.Struct {

			id := field.FieldByName("id")

			if id.IsValid() && id.Kind() == reflect.String {
				name := fieldType.Tag.Get("id")

				ptr := unsafe.Pointer(id.UnsafeAddr())

				if name == "" {
					name = fieldType.Name
				}

				reflect.NewAt(id.Type(), ptr).Elem().Set(reflect.ValueOf(name))
			}

		}
	}

	return *set
}

// BuildWithDefaults constructs an Options map by applying a list of Option[T] functions,
// optionally starting from a set of default values.
func BuildWithDefaults[T any](defaults Options, options ...Option[T]) Options {
	var opts Options

	if len(defaults) > 0 {
		opts = maps.Clone(defaults)
	} else {
		opts = make(Options)
	}

	for _, opt := range options {
		opt(opts)
	}

	return opts
}

// Build constructs an Options map from a variadic list of Option[T] functions.
func Build[T any](options ...Option[T]) Options {
	return BuildWithDefaults(nil, options...)
}
