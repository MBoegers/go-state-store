package storage

import "testing"

func TestSetString(t *testing.T) {
	var err = Set("testString", "test")
	if err != nil {
		t.Errorf("Unable to set String value, due to %v", err)
	}
}

func TestSetInt(t *testing.T) {
	var err = Set("testInt", 1)
	if err != nil {
		t.Errorf("Unable to set Integer value, due to %v", err)
	}
}
func TestSetBool(t *testing.T) {
	var err = Set("testBool", false)
	if err != nil {
		t.Errorf("Unable to set Boolean value, due to %v", err)
	}
}

func TestSetMultiple(t *testing.T) {
	Set("test", "value1")
	var err = Set("test", "value2")
	if err != nil {
		t.Errorf("Unable to set value twice, due to %v", err)
	}
}
