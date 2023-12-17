package main

import "testing"

func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {
		t.Errorf("expected bar, but got %s", result)
	}
}
