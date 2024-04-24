package test

import "testing"

func TestAdd(t *testing.T) {

	if Add(1, 2) == 3 {
		t.Log("1+2=3")
	}

	if Subtract(2, 3) == -1 {
		t.Log("2-2=-1")
	}
}
