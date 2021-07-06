package collections

import "testing"

func TestStringCreation(t *testing.T) {
	if CreateString() != "asd" {
		t.Errorf("not working\n")
	}
}
