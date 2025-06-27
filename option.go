package optkit

// Options is a map that holds option values.
type Options map[string]any

// Option is a function that modifies Options map.
type Option[T any] func(Options)

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

// GetWithDefault2 retrieves two values (T0, T1) from opts by ID.
// If the key doesn't exist or type assertion fails, it returns the provided defaults.
func GetWithDefault2[T0 any, T1 any](opts Options, id string, def0 T0, def1 T1) (T0, T1) {
	if v, ok := opts[id]; ok {
		if s, ok := v.([]any); ok && len(s) >= 2 {
			return s[0].(T0), s[1].(T1)
		}
	}
	return def0, def1
}

// Get2 retrieves two values (T0, T1) from opts by ID,
// or returns zero values if not present.
func Get2[T0 any, T1 any](opts Options, id string) (T0, T1) {
	return GetWithDefault2(opts, id, *new(T0), *new(T1))
}

// GetWithDefault3 retrieves three values from opts by ID, falling back to provided defaults.
func GetWithDefault3[T0, T1, T2 any](opts Options, id string, def0 T0, def1 T1, def2 T2) (T0, T1, T2) {
	if v, ok := opts[id]; ok {
		if s, ok := v.([]any); ok && len(s) >= 3 {
			return s[0].(T0), s[1].(T1), s[2].(T2)
		}
	}
	return def0, def1, def2
}

// Get3 fetches three values from opts or returns zero values.
func Get3[T0, T1, T2 any](opts Options, id string) (T0, T1, T2) {
	return GetWithDefault3(opts, id, *new(T0), *new(T1), *new(T2))
}

func GetWithDefault4[T0, T1, T2, T3 any](opts Options, id string, def0 T0, def1 T1, def2 T2, def3 T3) (T0, T1, T2, T3) {
	if v, ok := opts[id]; ok {
		if s, ok := v.([]any); ok && len(s) >= 4 {
			return s[0].(T0), s[1].(T1), s[2].(T2), s[3].(T3)
		}
	}
	return def0, def1, def2, def3
}

func Get4[T0, T1, T2, T3 any](opts Options, id string) (T0, T1, T2, T3) {
	return GetWithDefault4(opts, id, *new(T0), *new(T1), *new(T2), *new(T3))
}

func GetWithDefault5[T0, T1, T2, T3, T4 any](opts Options, id string, def0 T0, def1 T1, def2 T2, def3 T3, def4 T4) (T0, T1, T2, T3, T4) {
	if v, ok := opts[id]; ok {
		if s, ok := v.([]any); ok && len(s) >= 5 {
			return s[0].(T0), s[1].(T1), s[2].(T2), s[3].(T3), s[4].(T4)
		}
	}
	return def0, def1, def2, def3, def4
}

func Get5[T0, T1, T2, T3, T4 any](opts Options, id string) (T0, T1, T2, T3, T4) {
	return GetWithDefault5(opts, id, *new(T0), *new(T1), *new(T2), *new(T3), *new(T4))
}

func GetWithDefault6[T0, T1, T2, T3, T4, T5 any](opts Options, id string, def0 T0, def1 T1, def2 T2, def3 T3, def4 T4, def5 T5) (T0, T1, T2, T3, T4, T5) {
	if v, ok := opts[id]; ok {
		if s, ok := v.([]any); ok && len(s) >= 6 {
			return s[0].(T0), s[1].(T1), s[2].(T2), s[3].(T3), s[4].(T4), s[5].(T5)
		}
	}
	return def0, def1, def2, def3, def4, def5
}

func Get6[T0, T1, T2, T3, T4, T5 any](opts Options, id string) (T0, T1, T2, T3, T4, T5) {
	return GetWithDefault6(opts, id, *new(T0), *new(T1), *new(T2), *new(T3), *new(T4), *new(T5))
}
