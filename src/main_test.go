// Tests for main.go

package main

import "testing"

func TestSum(t *testing.T) {
	if add(3, 7) != 10 {
		t.Fail()
	}
	if add(-23, 77) != 54 {
		t.Fail()
	}
	if add(453, 107) != 560 {
		t.Fail()
	}
}

func TestProduct(t *testing.T) {
	if multiply(3, 7) != 21 {
		t.Fail()
	}
	if multiply(-6, 2) != -12 {
		t.Fail()
	}
	if multiply(4, 205) != 820 {
		t.Fail()
	}
}
