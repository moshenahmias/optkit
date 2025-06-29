package optkit_test

import (
	"testing"

	"github.com/moshenahmias/optkit"
)

func Test_Var_Set(t *testing.T) {
	const v0 optkit.Var[int, struct{}] = "v0"
	const v1 optkit.Var[string, struct{}] = "v1"

	opts := optkit.Build(v0.Set(42), v1.Set("Hello"), v1.Replace(func(s *string) {
		*s = *s + " World"
	}))

	if v := optkit.Get[int](opts, "v0"); v != 42 {
		t.Errorf("Expected 42, got %d", v)
	}

	if condition := optkit.Get[string](opts, "v1"); condition != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", condition)
	}
}

func Test_Var2_Set(t *testing.T) {
	const v optkit.Var2[string, int, struct{}] = "v2"

	opts := optkit.Build(
		v.Set("foo", 10),
		v.Replace(func(s *string, i *int) {
			*s += "-bar"
			*i *= 2
		}),
	)

	s, i := optkit.Get2[string, int](opts, "v2")
	if s != "foo-bar" || i != 20 {
		t.Errorf("Expected ('foo-bar', 20), got (%q, %d)", s, i)
	}
}

func Test_Var3_Set(t *testing.T) {
	const v optkit.Var3[int, int, int, struct{}] = "v3"

	opts := optkit.Build(
		v.Set(1, 2, 3),
		v.Replace(func(a, b, c *int) {
			*a += 10
			*b += 20
			*c += 30
		}),
	)

	a, b, c := optkit.Get3[int, int, int](opts, "v3")
	if a != 11 || b != 22 || c != 33 {
		t.Errorf("Expected (11, 22, 33), got (%d, %d, %d)", a, b, c)
	}
}

func Test_Var4_Set(t *testing.T) {
	const v optkit.Var4[string, string, int, bool, struct{}] = "v4"

	opts := optkit.Build(
		v.Set("a", "b", 5, false),
		v.Replace(func(a, b *string, i *int, flag *bool) {
			*a = *a + "x"
			*b = *b + "y"
			*i += 3
			*flag = true
		}),
	)

	a, b, i, flag := optkit.Get4[string, string, int, bool](opts, "v4")
	if a != "ax" || b != "by" || i != 8 || !flag {
		t.Errorf("Expected ('ax', 'by', 8, true), got (%q, %q, %d, %v)", a, b, i, flag)
	}
}

func Test_Var5_Set(t *testing.T) {
	const v optkit.Var5[int, int, int, int, int, struct{}] = "v5"

	opts := optkit.Build(
		v.Set(1, 2, 3, 4, 5),
		v.Replace(func(a, b, c, d, e *int) {
			*a *= 1
			*b *= 2
			*c *= 3
			*d *= 4
			*e *= 5
		}),
	)

	a, b, c, d, e := optkit.Get5[int, int, int, int, int](opts, "v5")
	if a != 1 || b != 4 || c != 9 || d != 16 || e != 25 {
		t.Errorf("Expected (1, 4, 9, 16, 25), got (%d, %d, %d, %d, %d)", a, b, c, d, e)
	}
}

func Test_Var6_Set(t *testing.T) {
	const v optkit.Var6[int, int, int, int, int, int, struct{}] = "v6"

	opts := optkit.Build(
		v.Set(1, 2, 3, 4, 5, 6),
		v.Replace(func(a, b, c, d, e, f *int) {
			*a += 1
			*b += 1
			*c += 1
			*d += 1
			*e += 1
			*f += 1
		}),
	)

	a, b, c, d, e, f := optkit.Get6[int, int, int, int, int, int](opts, "v6")
	if a != 2 || b != 3 || c != 4 || d != 5 || e != 6 || f != 7 {
		t.Errorf("Expected (2, 3, 4, 5, 6, 7), got (%d, %d, %d, %d, %d, %d)", a, b, c, d, e, f)
	}
}
