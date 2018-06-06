package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := strings.Repeat("x", bufio.MaxScanTokenSize)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}
