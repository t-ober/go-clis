package main

import (
	"bytes"
	"testing"
)

func TestCountBytes(t *testing.T) {
	r := bytes.NewBufferString("Hello")
	wc := count(r, false, true)
	if wc != 5 {
		t.Errorf("Expected 5, got %d", wc)
	}
}

func TestCountLines(t *testing.T) {
	r := bytes.NewBufferString("Hello \nWorld. Foo\n Bar")
	wc := count(r, true, false)
	if wc != 3 {
		t.Errorf("Expected 3, got %d", wc)
	}
}

func TestCountWords(t *testing.T) {
	r := bytes.NewBufferString("Hello World")
	wc := count(r, false, false)
	if wc != 2 {
		t.Errorf("Expected 2, got %d", wc)
	}
}
