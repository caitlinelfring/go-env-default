package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDefault(t *testing.T) {
	envVar := "MY_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetDefault(envVar, "default_value")
	assert.Equal(t, "default_value", val)

	os.Setenv(envVar, "defined_value")
	val = GetDefault(envVar, "default_value")
	assert.Equal(t, "defined_value", val)
}

func TestGetBoolDefault(t *testing.T) {
	envVar := "MY_BOOL_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetBoolDefault(envVar, true)
	assert.Equal(t, true, val)

	os.Setenv(envVar, "true")
	val = GetBoolDefault(envVar, false)
	assert.Equal(t, true, val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notABool")
	val = GetBoolDefault(envVar, true)
	assert.Equal(t, true, val)

	os.Unsetenv(envVar)
}

func TestGetIntDefault(t *testing.T) {
	envVar := "MY_INT_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetIntDefault(envVar, 11)
	assert.Equal(t, 11, val)

	os.Setenv(envVar, "1")
	val = GetIntDefault(envVar, 11)
	assert.Equal(t, 1, val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notAnInt")
	val = GetIntDefault(envVar, 11)
	assert.Equal(t, 11, val)
}

func TestGetInt64Default(t *testing.T) {
	envVar := "MY_INT_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetInt64Default(envVar, 11)
	assert.Equal(t, int64(11), val)

	os.Setenv(envVar, "1")
	val = GetInt64Default(envVar, 11)
	assert.Equal(t, int64(1), val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notAnInt")
	val = GetInt64Default(envVar, 11)
	assert.Equal(t, int64(11), val)
}

func TestGetFloatDefault(t *testing.T) {
	envVar := "MY_FLOAT_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetFloatDefault(envVar, 1.1)
	assert.Equal(t, 1.1, val)

	os.Setenv(envVar, "1.2")
	val = GetFloatDefault(envVar, 1.1)
	assert.Equal(t, 1.2, val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notAnInt")
	val = GetFloatDefault(envVar, 1.1)
	assert.Equal(t, 1.1, val)
}

func TestGetDurationDefault(t *testing.T) {
	envVar := "MY_DUR_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})
	val := GetDurationDefault(envVar, time.Hour)
	assert.Equal(t, time.Hour, val)

	os.Setenv(envVar, "1m")
	val = GetDurationDefault(envVar, time.Hour)
	assert.Equal(t, time.Minute, val)

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notADuration")
	val = GetDurationDefault(envVar, time.Hour)
	assert.Equal(t, time.Hour, val)
}
