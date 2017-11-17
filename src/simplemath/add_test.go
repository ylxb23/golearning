package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(2, 3)
	if r != 5 {
		t.Errorf("Add(2, 3) failed. Got %d, expected 5", r);
	}
}