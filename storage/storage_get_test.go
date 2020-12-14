package storage

import "testing"

func TestGetEmpty(t *testing.T) {
	var actual = Get("test")
	if actual != nil {
		t.Errorf("should read nil on empty store")
	}
}

func TestGetNotSet(t *testing.T) {
	var key = "test"
	var otherKey = "otherTest"
	var value = "value"
	Set(key, value)
	var expected = Get(otherKey)
	if expected != nil {
		t.Errorf("should read nil from empty store, got %v", value)
	}
}

func TestGetString(t *testing.T) {
	var key = "test"
	var value = "value"
	Set(key, value)
	var expected = Get(key)
	if expected != value {
		t.Errorf("should read %v from empty store, got %v", value, expected)
	}
}
