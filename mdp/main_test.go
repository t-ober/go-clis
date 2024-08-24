package main

import (
	"bytes"
	"os"
	"regexp"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	re := regexp.MustCompile(`\s+`)
	expectedClean := re.ReplaceAll(expected, []byte{})
	resultClean := re.ReplaceAll(result, []byte{})

	if !bytes.Equal(expectedClean, resultClean) {
		t.Logf("golden:\n%s\n", expectedClean)
		t.Logf("result:\n%s\n", resultClean)
		t.Error("Result content does not match the golden file")
	}

}

func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer

	if err := run(inputFile, "", &mockStdOut, true); err != nil {
		t.Fatal(err)
	}

	resultFile := strings.TrimSpace(mockStdOut.String())
	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	re := regexp.MustCompile(`\s+`)
	expectedClean := re.ReplaceAll(expected, []byte{})
	resultClean := re.ReplaceAll(result, []byte{})

	if !bytes.Equal(expectedClean, resultClean) {
		t.Logf("golden:\n%s\n", expectedClean)
		t.Logf("result:\n%s\n", resultClean)
		t.Error("Result content does not match the golden file")
	}

	os.Remove(resultFile)
}
