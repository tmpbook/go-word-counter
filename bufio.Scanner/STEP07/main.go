package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	if err == nil && token != nil && bytes.Equal(token, []byte{'e', 'n', 'd'}) {
		return 0, []byte{'E', 'N', 'D'}, bufio.ErrFinalToken
	}
	return
}
func main() {
	input := "foo end bar"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(split)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Printf("Error: %s\n", scanner.Err())
	}
}
