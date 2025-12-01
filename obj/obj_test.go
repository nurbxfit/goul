package obj

import "testing"

func TestKeys(t *testing.T) {
	type TestStrucct struct {
		Name string
		Age  int
	}

	keys := Keys(TestStrucct{Name: "Nur", Age: 30})

	if len(keys) != 2 {
		t.Errorf("Expected 2 keys got %d", len(keys))
	}

	if keys[0] != "Name" {
		t.Errorf("Expected first key to be Name, got %s", keys[0])
	}
}
