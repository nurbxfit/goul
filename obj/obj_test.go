package obj

import "testing"

type TestStruct struct {
	Name string
	Age  int
}

func TestKeys(t *testing.T) {

	keys := Keys(TestStruct{Name: "Nur", Age: 30})

	if len(keys) != 2 {
		t.Errorf("Expected 2 keys got %d", len(keys))
	}

	if keys[0] != "Name" {
		t.Errorf("Expected first key to be Name, got %s", keys[0])
	}
}

func TestValues(t *testing.T) {
	values := Values(TestStruct{Name: "Nur", Age: 30})

	if len(values) != 2 {
		t.Errorf("Expected 2 values got %d", len(values))
	}

	if values[0] != "Nur" {
		t.Errorf("Expected first key to be Nur, got %s", values[0])
	}
}

func TestEntries(t *testing.T) {
	entries := Entries(TestStruct{Name: "Nur", Age: 30})
	if len(entries) != 2 {
		t.Errorf("Expected 2 values got %d", len(entries))
	}

	// expected entries
	expected := []Entry{
		{Key: "Name", Value: "Nur"},
		{Key: "Age", Value: 30},
	}

	for i, exp := range expected {
		if entries[i].Key != exp.Key {
			t.Errorf("Entry %d: expected key %s, got %s", i, exp.Key, entries[i].Key)
		}

		if entries[i].Value != exp.Value {
			t.Errorf("Entry %d: expected value %v, got %v", i, exp.Value, entries[i].Value)
		}
	}
}
