package env

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func assertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Not equal: \n"+
			"expected: %v\n"+
			"actual  : %v", expected, actual)
	}
}

func TestGetDefault(t *testing.T) {
	envVar := "MY_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetDefault(envVar, "default_value")
	assertEqual(t, "default_value", val)

	os.Setenv(envVar, "defined_value")
	val = GetDefault(envVar, "default_value")
	assertEqual(t, "defined_value", val)
}

func TestGetBoolDefault(t *testing.T) {
	envVar := "MY_BOOL_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetBoolDefault(envVar, true)
	assertEqual(t, true, val)

	os.Setenv(envVar, "true")
	val = GetBoolDefault(envVar, false)
	assertEqual(t, true, val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notABool")
	val = GetBoolDefault(envVar, true)
	assertEqual(t, true, val)

	os.Unsetenv(envVar)
}

func TestGetIntDefault(t *testing.T) {
	envVar := "MY_INT_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetIntDefault(envVar, 11)
	assertEqual(t, 11, val)

	os.Setenv(envVar, "1")
	val = GetIntDefault(envVar, 11)
	assertEqual(t, 1, val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notAnInt")
	val = GetIntDefault(envVar, 11)
	assertEqual(t, 11, val)
}

func TestGetInt64Default(t *testing.T) {
	envVar := "MY_INT_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetInt64Default(envVar, 11)
	assertEqual(t, int64(11), val)

	os.Setenv(envVar, "1")
	val = GetInt64Default(envVar, 11)
	assertEqual(t, int64(1), val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notAnInt")
	val = GetInt64Default(envVar, 11)
	assertEqual(t, int64(11), val)
}

func TestGetFloatDefault(t *testing.T) {
	envVar := "MY_FLOAT_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetFloatDefault(envVar, 1.1)
	assertEqual(t, 1.1, val)

	os.Setenv(envVar, "1.2")
	val = GetFloatDefault(envVar, 1.1)
	assertEqual(t, 1.2, val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notAnInt")
	val = GetFloatDefault(envVar, 1.1)
	assertEqual(t, 1.1, val)
}

func TestGetDurationDefault(t *testing.T) {
	envVar := "MY_DUR_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetDurationDefault(envVar, time.Hour)
	assertEqual(t, time.Hour, val)

	os.Setenv(envVar, "1m")
	val = GetDurationDefault(envVar, time.Hour)
	assertEqual(t, time.Minute, val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notADuration")
	val = GetDurationDefault(envVar, time.Hour)
	assertEqual(t, time.Hour, val)
}

func ExampleGetDefault() {
	fmt.Println(GetDefault("MY_ENV_VAR", "default"))
	os.Setenv("MY_ENV_VAR", "custom")
	fmt.Println(GetDefault("MY_ENV_VAR", "default"))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// default
	// custom
}
func ExampleGetBoolDefault() {
	fmt.Println(GetBoolDefault("MY_ENV_VAR", true))
	os.Setenv("MY_ENV_VAR", "false")
	fmt.Println(GetBoolDefault("MY_ENV_VAR", true))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// true
	// false
}
func ExampleGetIntDefault() {
	fmt.Println(GetIntDefault("MY_ENV_VAR", 1))
	os.Setenv("MY_ENV_VAR", "100")
	fmt.Println(GetIntDefault("MY_ENV_VAR", 1))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// 1
	// 100

}
func ExampleGetInt64Default() {
	fmt.Println(GetInt64Default("MY_ENV_VAR", int64(12345678910)))
	os.Setenv("MY_ENV_VAR", "100")
	fmt.Println(GetInt64Default("MY_ENV_VAR", int64(12345678910)))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// 12345678910
	// 100
}
func ExampleGetFloatDefault() {
	fmt.Println(GetFloatDefault("MY_ENV_VAR", 3.14))
	os.Setenv("MY_ENV_VAR", "34.02")
	fmt.Println(GetFloatDefault("MY_ENV_VAR", 3.14))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// 3.14
	// 34.02
}
func ExampleGetDurationDefault() {
	fmt.Println(GetDurationDefault("MY_ENV_VAR", time.Hour))
	os.Setenv("MY_ENV_VAR", "60s")
	fmt.Println(GetDurationDefault("MY_ENV_VAR", time.Hour))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// 1h0m0s
	// 1m0s
}
