package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hellos := []*String{
		&String{Value: "Hello"},
		&String{Value: "Hello"},
	}

	goodbyes := []*String{
		&String{Value: "goodbye"},
		&String{Value: "goodbye"},
	}

	if hellos[0].HashKey() != hellos[1].HashKey() {
		t.Error("strings with same content have different hash keys")
	}

	if goodbyes[0].HashKey() != goodbyes[1].HashKey() {
		t.Error("strings with same content have different hash keys")
	}

	if hellos[0].HashKey() == goodbyes[0].HashKey() {
		t.Error("strings with different content have same hash keys")
	}
}
