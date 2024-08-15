package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Printf("Count: %d", count(os.Stdin, *lines, *bytes))
}

func count(reader io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(reader)
	if countLines {
		scanner.Split(bufio.ScanLines)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
