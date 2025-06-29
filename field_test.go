package optkit_test

import (
	"testing"

	"github.com/moshenahmias/optkit"
)

func Test_Set(t *testing.T) {
	type TestOptions struct {
		Field1 optkit.Field[int, TestOptions]
		Field2 optkit.Field[string, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.Build(options.Field1.Set(42), options.Field2.Set("Hello"))

	if v := optkit.Get[int](opts, "Field1"); v != 42 {
		t.Errorf("Expected 42, got %d", v)
	}

	if condition := optkit.Get[string](opts, "Field2"); condition != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", condition)
	}
}

func Test_Replace(t *testing.T) {
	type TestOptions struct {
		Field1 optkit.Field[int, TestOptions]
		Field2 optkit.Field[string, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.Build(options.Field1.Set(42), options.Field2.Set("Hello"), options.Field1.Replace(func(v *int) {
		*v = *v + 10
	}))

	if v := optkit.Get[int](opts, "Field1"); v != 52 {
		t.Errorf("Expected 52, got %d", v)
	}

	if condition := optkit.Get[string](opts, "Field2"); condition != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", condition)
	}
}

func Test_IdTag(t *testing.T) {
	type TestOptions struct {
		Field1 optkit.Field[int, TestOptions] `id:"f1"`
		Field2 optkit.Field[string, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.Build(options.Field1.Set(42), options.Field2.Set("Hello"))

	if v := optkit.Get[int](opts, "f1"); v != 42 {
		t.Errorf("Expected 42, got %d", v)
	}

	if condition := optkit.Get[string](opts, "Field2"); condition != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", condition)
	}
}

func Test_Defaults(t *testing.T) {
	type TestOptions struct {
		Field1 optkit.Field[int, TestOptions]
		Field2 optkit.Field[string, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.BuildWithDefaults(optkit.Options{"Field2": "Hello"}, options.Field1.Set(42))

	if v := optkit.Get[int](opts, "Field1"); v != 42 {
		t.Errorf("Expected 42, got %d", v)
	}

	if condition := optkit.Get[string](opts, "Field2"); condition != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", condition)
	}
}

func Test_Field2(t *testing.T) {
	type TestOptions struct {
		Field optkit.Field2[int, string, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.Build(options.Field.Set(3, "Hello"))

	if v0, v1 := optkit.Get2[int, string](opts, "Field"); v0 != 3 || v1 != "Hello" {
		t.Errorf("Expected (3, 'Hello'), got (%d, %s)", v0, v1)
	}
}

func Test_Field3(t *testing.T) {
	type TestOptions struct {
		Field optkit.Field3[int, string, bool, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.Build(options.Field.Set(7, "World", true))

	if v0, v1, v2 := optkit.Get3[int, string, bool](opts, "Field"); v0 != 7 || v1 != "World" || v2 != true {
		t.Errorf("Expected (7, 'World', true), got (%d, %s, %v)", v0, v1, v2)
	}
}

func Test_Field4(t *testing.T) {
	type TestOptions struct {
		Field optkit.Field4[int, string, bool, float64, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.Build(options.Field.Set(1, "Go", false, 3.14))

	if v0, v1, v2, v3 := optkit.Get4[int, string, bool, float64](opts, "Field"); v0 != 1 || v1 != "Go" || v2 != false || v3 != 3.14 {
		t.Errorf("Expected (1, 'Go', false, 3.14), got (%d, %s, %v, %f)", v0, v1, v2, v3)
	}
}

func Test_Field5(t *testing.T) {
	type TestOptions struct {
		Field optkit.Field5[int, string, bool, float64, byte, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.Build(options.Field.Set(9, "Kit", true, 2.71, 'A'))

	if v0, v1, v2, v3, v4 := optkit.Get5[int, string, bool, float64, byte](opts, "Field"); v0 != 9 || v1 != "Kit" || v2 != true || v3 != 2.71 || v4 != 'A' {
		t.Errorf("Expected (9, 'Kit', true, 2.71, 'A'), got (%d, %s, %v, %f, %c)", v0, v1, v2, v3, v4)
	}
}

func Test_Field6(t *testing.T) {
	type TestOptions struct {
		Field optkit.Field6[int, string, bool, float64, byte, rune, TestOptions]
	}

	options := optkit.Init[TestOptions]()

	opts := optkit.Build(options.Field.Set(5, "Zebra", false, 1.23, 'X', 'λ'))

	if v0, v1, v2, v3, v4, v5 := optkit.Get6[int, string, bool, float64, byte, rune](opts, "Field"); v0 != 5 || v1 != "Zebra" || v2 != false || v3 != 1.23 || v4 != 'X' || v5 != 'λ' {
		t.Errorf("Expected (5, 'Zebra', false, 1.23, 'X', 'λ'), got (%d, %s, %v, %f, %c, %c)", v0, v1, v2, v3, v4, v5)
	}
}
