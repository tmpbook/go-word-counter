package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "foo\nbar\nbaz"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Not actually needed since it’s a default split function.
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
