package main

import (
	"testing"
)

func TestCreateMessage(t *testing.T) {
	expected := "{\"myName\":\"piem\",\"cputemp\":37,\"cpuload\":2,\"sine\":0.9}"
	result := string(createMessage())
	if result != expected {
		t.Fatalf("expected %s, got %s.", expected, result)
	}
}
