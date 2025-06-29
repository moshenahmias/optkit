package optkit

// Var is a typed key for a single value of type T stored in a map of options.
type Var[T any, S any] string

// Set assigns a value of type T to the options map under this Var key.
func (v Var[T, S]) Set(t T) Option[S] {
	return Option[S](func(m Options) {
		m[string(v)] = t
	})
}

// Replace retrieves the current value using Get, allows mutation via the provided function,
// then updates the modified value back into the options map.
func (v Var[T, S]) Replace(fn func(*T)) Option[S] {
	return Option[S](func(m Options) {
		t := Get[T](m, string(v)) // Get returns zero value if missing or wrong type
		fn(&t)
		m[string(v)] = t
	})
}

// Var2 stores two typed values.
type Var2[T0 any, T1 any, S any] string

func (v Var2[T0, T1, S]) Set(t0 T0, t1 T1) Option[S] {
	return Option[S](func(m Options) {
		m[string(v)] = []any{t0, t1}
	})
}

// Replace uses Get2 to retrieve the two values (or zero values), mutates them in-place,
// and stores them back into the options map as a slice.
func (v Var2[T0, T1, S]) Replace(fn func(*T0, *T1)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1 := Get2[T0, T1](m, string(v)) // Get2 returns zero values if missing/malformed
		fn(&v0, &v1)
		m[string(v)] = []any{v0, v1}
	})
}

// Var3 stores three typed values.
type Var3[T0 any, T1 any, T2 any, S any] string

func (v Var3[T0, T1, T2, S]) Set(t0 T0, t1 T1, t2 T2) Option[S] {
	return Option[S](func(m Options) {
		m[string(v)] = []any{t0, t1, t2}
	})
}

// Replace retrieves the tuple using Get3, applies the mutation, and stores it back.
func (v Var3[T0, T1, T2, S]) Replace(fn func(*T0, *T1, *T2)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1, v2 := Get3[T0, T1, T2](m, string(v))
		fn(&v0, &v1, &v2)
		m[string(v)] = []any{v0, v1, v2}
	})
}

// Var4 stores four typed values.
type Var4[T0 any, T1 any, T2 any, T3 any, S any] string

func (v Var4[T0, T1, T2, T3, S]) Set(t0 T0, t1 T1, t2 T2, t3 T3) Option[S] {
	return Option[S](func(m Options) {
		m[string(v)] = []any{t0, t1, t2, t3}
	})
}

// Replace retrieves the values via Get4, updates via fn, and stores them again.
func (v Var4[T0, T1, T2, T3, S]) Replace(fn func(*T0, *T1, *T2, *T3)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1, v2, v3 := Get4[T0, T1, T2, T3](m, string(v))
		fn(&v0, &v1, &v2, &v3)
		m[string(v)] = []any{v0, v1, v2, v3}
	})
}

// Var5 stores five typed values.
type Var5[T0 any, T1 any, T2 any, T3 any, T4 any, S any] string

func (v Var5[T0, T1, T2, T3, T4, S]) Set(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4) Option[S] {
	return Option[S](func(m Options) {
		m[string(v)] = []any{t0, t1, t2, t3, t4}
	})
}

// Replace retrieves five values using Get5, mutates them, and reassigns them to the map.
func (v Var5[T0, T1, T2, T3, T4, S]) Replace(fn func(*T0, *T1, *T2, *T3, *T4)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1, v2, v3, v4 := Get5[T0, T1, T2, T3, T4](m, string(v))
		fn(&v0, &v1, &v2, &v3, &v4)
		m[string(v)] = []any{v0, v1, v2, v3, v4}
	})
}

// Var6 stores six typed values.
type Var6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, S any] string

func (v Var6[T0, T1, T2, T3, T4, T5, S]) Set(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) Option[S] {
	return Option[S](func(m Options) {
		m[string(v)] = []any{t0, t1, t2, t3, t4, t5}
	})
}

// Replace uses Get6 to retrieve and mutate six typed values and store them in the options map.
func (v Var6[T0, T1, T2, T3, T4, T5, S]) Replace(fn func(*T0, *T1, *T2, *T3, *T4, *T5)) Option[S] {
	return Option[S](func(m Options) {
		v0, v1, v2, v3, v4, v5 := Get6[T0, T1, T2, T3, T4, T5](m, string(v))
		fn(&v0, &v1, &v2, &v3, &v4, &v5)
		m[string(v)] = []any{v0, v1, v2, v3, v4, v5}
	})
}
