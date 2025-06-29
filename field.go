package optkit

// Field represents a field in an options struct.
type Field[T any, S any] struct {
	id string
}

// Set sets the value of the field in the Options map.
func (f Field[T, S]) Set(v T) Option[S] {
	return Option[S](func(m Options) {
		m[f.id] = v
	})
}

// Replace applies a function to the current value of the field in the Options map.
func (f Field[T, S]) Replace(fn func(*T)) Option[S] {
	return Option[S](func(m Options) {
		v := Get[T](m, f.id)
		fn(&v)
		m[f.id] = v
	})
}

// Field2 represents an option field storing two typed values (T0 and T1)
// associated with a parent struct type S. The 'id' uniquely identifies the field.
type Field2[T0 any, T1 any, S any] struct {
	id string
}

// Set stores values v0 and v1 into the options map under this field's ID.
func (f Field2[T0, T1, S]) Set(v0 T0, v1 T1) Option[S] {
	return Option[S](func(m Options) {
		m[f.id] = []any{v0, v1}
	})
}

// Replace retrieves the stored values, lets the caller mutate them,
// and updates the map with the modified values.
func (f Field2[T0, T1, S]) Replace(fn func(*T0, *T1)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1 := Get2[T0, T1](m, f.id)
		fn(&v0, &v1)
		m[f.id] = []any{v0, v1}
	})
}

// Field3 represents a field that stores three typed values (T0, T1, T2)
// bound to an owner struct type S.
type Field3[T0 any, T1 any, T2 any, S any] struct {
	id string
}

// Set inserts all three values into the options map.
func (f Field3[T0, T1, T2, S]) Set(v0 T0, v1 T1, v2 T2) Option[S] {
	return Option[S](func(m Options) {
		m[f.id] = []any{v0, v1, v2}
	})
}

// Replace retrieves the three values, applies user-supplied mutation logic,
// then updates the options map with the new values.
func (f Field3[T0, T1, T2, S]) Replace(fn func(*T0, *T1, *T2)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1, v2 := Get3[T0, T1, T2](m, f.id)
		fn(&v0, &v1, &v2)
		m[f.id] = []any{v0, v1, v2}
	})
}

// Field4 defines a key for storing four typed values (T0, T1, T2, T3)
// under an ID in an options map associated with struct S.
type Field4[T0, T1, T2, T3, S any] struct {
	id string
}

// Set stores four values into the options map for this field.
func (f Field4[T0, T1, T2, T3, S]) Set(v0 T0, v1 T1, v2 T2, v3 T3) Option[S] {
	return Option[S](func(m Options) {
		m[f.id] = []any{v0, v1, v2, v3}
	})
}

// Replace allows in-place modification of the stored values using the provided function.
func (f Field4[T0, T1, T2, T3, S]) Replace(fn func(*T0, *T1, *T2, *T3)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1, v2, v3 := Get4[T0, T1, T2, T3](m, f.id)
		fn(&v0, &v1, &v2, &v3)
		m[f.id] = []any{v0, v1, v2, v3}
	})
}

// Field5 defines a field for five typed values (T0 through T4), stored in Options under ID.
type Field5[T0, T1, T2, T3, T4, S any] struct {
	id string
}

// Set places five values into the options map for this field key.
func (f Field5[T0, T1, T2, T3, T4, S]) Set(v0 T0, v1 T1, v2 T2, v3 T3, v4 T4) Option[S] {
	return Option[S](func(m Options) {
		m[f.id] = []any{v0, v1, v2, v3, v4}
	})
}

// Replace lets the user mutate the stored five values directly.
func (f Field5[T0, T1, T2, T3, T4, S]) Replace(fn func(*T0, *T1, *T2, *T3, *T4)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1, v2, v3, v4 := Get5[T0, T1, T2, T3, T4](m, f.id)
		fn(&v0, &v1, &v2, &v3, &v4)
		m[f.id] = []any{v0, v1, v2, v3, v4}
	})
}

// Field6 lets you define a field that stores six typed values in Options under a specific ID.
type Field6[T0, T1, T2, T3, T4, T5, S any] struct {
	id string
}

// Set inserts six values into the options map with this field's ID.
func (f Field6[T0, T1, T2, T3, T4, T5, S]) Set(v0 T0, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Option[S] {
	return Option[S](func(m Options) {
		m[f.id] = []any{v0, v1, v2, v3, v4, v5}
	})
}

// Replace allows modification of the six stored values using the supplied function.
func (f Field6[T0, T1, T2, T3, T4, T5, S]) Replace(fn func(*T0, *T1, *T2, *T3, *T4, *T5)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1, v2, v3, v4, v5 := Get6[T0, T1, T2, T3, T4, T5](m, f.id)
		fn(&v0, &v1, &v2, &v3, &v4, &v5)
		m[f.id] = []any{v0, v1, v2, v3, v4, v5}
	})
}
