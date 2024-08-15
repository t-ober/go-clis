package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Printf("Count: %d", count(os.Stdin))
}

func count(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
