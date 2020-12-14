package storage

import (
	"testing"
)

func TestRemoveEmpty(t *testing.T) {
	var err = Remove("testEmpty")
	if err == nil {
		t.Errorf("should get an error removing not existing value")
	}
}

func TestRemoveSimple(t *testing.T) {
	var key = "testSimple"
	var value = "testValue"
	Set(key, value)

	var firstRead = Get(key)
	Remove(key)
	var secondRead = Get(key)

	if firstRead != value {
		t.Errorf("value should be present bevor deletion")
	}
	if secondRead != nil {
		t.Errorf("value should be absend after deletion")
	}
}
