package goenv

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}
}

// GetString returns the value of the environment variable with the given key.
// If the variable is not set, it returns the defaultValue.
func GetString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

// GetInt returns the value of the environment variable with the given key as an integer.
// If the variable is not set, it returns the defaultValue.
// If the value cannot be converted to an integer, it returns an error.
func GetInt(key string, defaultValue int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetInt64 returns the value of the environment variable with the given key as a 64-bit integer.
// If the variable is not set, it returns the defaultValue.
// If the value cannot be converted to a 64-bit integer, it returns an error.
func GetInt64(key string, defaultValue int64) (int64, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetFloat64 returns the value of the environment variable with the given key as a 64-bit floating-point number.
// If the variable is not set, it returns the defaultValue.
// If the value cannot be converted to a 64-bit floating-point number, it returns an error.
func GetFloat64(key string, defaultValue float64) (float64, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// GetBool returns the value of the environment variable with the given key as a boolean.
// If the variable is not set, it returns the defaultValue.
// If the value is not "true" or "false" (ignoring case), it returns an error.
func GetBool(key string, defaultValue bool) (bool, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return b, nil
}
