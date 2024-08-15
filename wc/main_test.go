package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	r := bytes.NewBufferString("Hello World")
	wc := count(r)
	if wc != 2 {
		t.Errorf("Expected 2, got %d", wc)
	}
}
