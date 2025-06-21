package optkit

import (
	"maps"
	"reflect"
	"unsafe"
)

// Field represents a field in an options struct.
type Field[T any, S any] struct {
	id string
}

// Options is a map that holds option values.
type Options map[string]any

// GetWithDefault retrieves the value associated with the given `id` from the Options map.
// If the value is not found or cannot be cast to the desired type, it returns the provided default value.
func GetWithDefault[T any](opts Options, id string, def T) T {
	if v, ok := opts[id]; ok {
		if t, ok := v.(T); ok {
			return t
		}
	}
	return def
}

// Get retrieves the value associated with the given `id` from the Options map.
// If the value is not found or cannot be cast to the desired type, it returns the zero value of type T.
func Get[T any](opts Options, id string) T {
	return GetWithDefault(opts, id, *new(T))
}

// Option is a function that modifies Options map.
type Option[T any] func(Options)

// Set sets the value of the field in the Options map.
func (f Field[T, S]) Set(v T) Option[S] {
	return Option[S](func(m Options) {
		m[f.id] = v
	})
}

// Replace applies a function to the current value of the field in the Options map.
func (f Field[T, S]) Replace(fn func(T) T) Option[S] {
	return Option[S](func(m Options) {
		m[f.id] = fn(Get[T](m, f.id))
	})
}

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
