package goenv

import (
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	key := "MY_VAR"
	defaultValue := "default"

	// Test default value
	value := GetString(key, defaultValue)
	if value != defaultValue {
		t.Errorf("expected %s, got %s", defaultValue, value)
	}

	// Test set value
	os.Setenv(key, "test")
	defer os.Unsetenv(key)
	value = GetString(key, defaultValue)
	if value != "test" {
		t.Errorf("expected %s, got %s", "test", value)
	}
}

func TestGetInt(t *testing.T) {
	key := "MY_VAR"
	defaultValue := 42

	// Test default value
	value, err := GetInt(key, defaultValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != defaultValue {
		t.Errorf("expected %d, got %d", defaultValue, value)
	}

	// Test set value
	os.Setenv(key, "24")
	defer os.Unsetenv(key)
	value, err = GetInt(key, defaultValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != 24 {
		t.Errorf("expected %d, got %d", 24, value)
	}

	// Test invalid value
	os.Setenv(key, "not an int")
	defer os.Unsetenv(key)
	_, err = GetInt(key, defaultValue)
	if err == nil {
		t.Errorf("expected error, got no error")
	}
}

func TestGetInt64(t *testing.T) {
	key := "MY_VAR"
	defaultValue := int64(42)

	// Test default value
	value, err := GetInt64(key, defaultValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != defaultValue {
		t.Errorf("expected %d, got %d", defaultValue, value)
	}

	// Test set value
	os.Setenv(key, "24")
	defer os.Unsetenv(key)
	value, err = GetInt64(key, defaultValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != 24 {
		t.Errorf("expected %d, got %d", 24, value)
	}

	// Test invalid value
	os.Setenv(key, "not an int")
	defer os.Unsetenv(key)
	_, err = GetInt64(key, defaultValue)
	if err == nil {
		t.Errorf("expected error, got no error")
	}
}

func TestGetFloat64(t *testing.T) {
	key := "MY_VAR"
	defaultValue := 3.14

	// Test default value
	value, err := GetFloat64(key, defaultValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != defaultValue {
		t.Errorf("expected %f, got %f", defaultValue, value)
	}

	// Test set value
	os.Setenv(key, "2.71")
	defer os.Unsetenv(key)
	value, err = GetFloat64(key, defaultValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != 2.71 {
		t.Errorf("expected %f, got %f", 2.71, value)
	}

	// Test invalid value
	os.Setenv(key, "not a float")
	defer os.Unsetenv(key)
	_, err = GetFloat64(key, defaultValue)
	if err == nil {
		t.Errorf("expected error, got no error")
	}
}

func TestGetBool(t *testing.T) {
	key := "MY_VAR"
	defaultValue := true

	// Test default value
	value, err := GetBool(key, defaultValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != defaultValue {
		t.Errorf("expected %t, got %t", defaultValue, value)
	}

	// Test set value
	os.Setenv(key, "true")
	defer os.Unsetenv(key)
	value, err = GetBool(key, defaultValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != true {
		t.Errorf("expected %t, got %t", true, value)
	}

	// Test invalid value
	os.Setenv(key, "not a bool")
	defer os.Unsetenv(key)
	_, err = GetBool(key, defaultValue)
	if err == nil {
		t.Errorf("expected error, got no error")
	}
}
