package datastore

import "testing"

func TestStorage(t *testing.T) {
	beforeAll()
	t.Run("Test set function", TestSet)
	t.Run("Test remove function", TestRemove)
	t.Run("Test gets function", TestGet)
}

func beforeAll() {
	var channel = make(chan string)
	Init(channel)
	go func() {
		for {
			<-channel
		}
	}()
}

//region<Set>
func TestSet(t *testing.T) {
	t.Run("Test String value", setString)
	t.Run("Test Intege value", setInt)
	t.Run("Test Boolean value", setBool)
	t.Run("Test multiple times", setMultiple)
}

func setString(t *testing.T) {
	var err = Set("testSetString", "test")
	if err != nil {
		t.Errorf("Unable to set String value, due to %v", err)
	}
}

func setInt(t *testing.T) {
	var err = Set("testSetInt", 1)
	if err != nil {
		t.Errorf("Unable to set Integer value, due to %v", err)
	}
}
func setBool(t *testing.T) {
	var err = Set("testSetBool", false)
	if err != nil {
		t.Errorf("Unable to set Boolean value, due to %v", err)
	}
}

func setMultiple(t *testing.T) {
	Set("testSetMulti", "value1")
	var err = Set("testSetMulti", "value2")
	if err != nil {
		t.Errorf("Unable to set value twice, due to %v", err)
	}
}

//endregion

//region<Remove>
func TestRemove(t *testing.T) {
	t.Run("Remove from empty", removeEmpty)
	t.Run("Remove simple", removeSimple)
}

func removeEmpty(t *testing.T) {
	var err = Remove("testRemoveEmpty")
	if err == nil {
		t.Errorf("should get an error removing not existing value")
	}
}

func removeSimple(t *testing.T) {
	var key = "testRemoveSimple"
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

//endregion

//region<Get>
func TestGet(t *testing.T) {
	t.Run("Get from empty", getEmpty)
	t.Run("Get not set value", getNotSet)
	t.Run("Get string value", getString)
}

func getEmpty(t *testing.T) {
	var actual = Get("testGetEmpty")
	if actual != nil {
		t.Errorf("should read nil on empty store")
	}
}

func getNotSet(t *testing.T) {
	var key = "testGetNotSet"
	var otherKey = "otherTestGet"
	var value = "value"
	Set(key, value)
	var expected = Get(otherKey)
	if expected != nil {
		t.Errorf("should read nil from empty store, got %v", value)
	}
}

func getString(t *testing.T) {
	var key = "testGetString"
	var value = "value"
	Set(key, value)
	var expected = Get(key)
	if expected != value {
		t.Errorf("should read %v from empty store, got %v", value, expected)
	}
}

//endregion
