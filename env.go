package env

import (
	"os"
	"strconv"
	"time"
)

// GetDefault returns the string value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetDefault(envVar, defaultValue string) string {
	if v, ok := os.LookupEnv(envVar); ok && len(v) > 0 {
		return v
	}
	return defaultValue
}

// GetBoolDefault returns the boolean value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetBoolDefault(envVar string, defaultValue bool) bool {
	val := GetDefault(envVar, strconv.FormatBool(defaultValue))
	if b, err := strconv.ParseBool(val); err == nil {
		return b
	}
	return defaultValue
}

// GetIntDefault returns the int value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetIntDefault(envVar string, defaultValue int) int {
	val := GetDefault(envVar, strconv.Itoa(defaultValue))
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}
	return defaultValue
}

// GetInt64Default returns the int64 value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetInt64Default(envVar string, defaultValue int64) int64 {
	val := GetDefault(envVar, strconv.FormatInt(defaultValue, 16))
	if i, err := strconv.ParseInt(val, 10, 64); err == nil {
		return i
	}
	return defaultValue
}

// GetFloatDefault returns the float64 value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetFloatDefault(envVar string, defaultValue float64) float64 {
	val := GetDefault(envVar, strconv.FormatFloat(defaultValue, 'E', -1, 64))
	if f, err := strconv.ParseFloat(val, 64); err == nil {
		return f
	}
	return defaultValue
}

// GetDurationDefault returns the time.Duration value of the environment variable, or a default
// value if the environment variable is not defined or is an empty string
func GetDurationDefault(envVar string, defaultValue time.Duration) time.Duration {
	val := GetDefault(envVar, defaultValue.String())
	if t, err := time.ParseDuration(val); err == nil {
		return t
	}
	return defaultValue
}
