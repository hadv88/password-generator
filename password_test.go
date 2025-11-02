package main

import "testing"

func TestGeneratePassword(t *testing.T) {
	pass := generatePassword(10, true)
	if len(pass) != 10 {
		t.Errorf("Expected length 10, got %d", len(pass))
	}
}
