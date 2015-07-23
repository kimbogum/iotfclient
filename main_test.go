package main

import (
	"testing"
)

func TestCreateMessage(t *testing.T) {
	expected := "{\"Temp\":56}"
	result := string(createMessage())
	if  result !=  expected {
		t.Fatalf("expected %s, got %s.", expected, result )
	}
}