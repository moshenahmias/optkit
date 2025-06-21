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

	opts := optkit.Build(options.Field1.Set(42), options.Field2.Set("Hello"), options.Field1.Replace(func(v int) int {
		return v + 10
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
