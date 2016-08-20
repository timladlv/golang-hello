package main

import "testing"

func TestMessage(t *testing.T) {
	if message() != "Hello" {
		t.Error("does not equal Hello")
	} else {
		// passed
	}
}
