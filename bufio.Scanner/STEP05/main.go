package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := "foofoofoo"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if bytes.Equal(data[:3], []byte{'f', 'o', 'o'}) {
			return 3, []byte{'F'}, nil
		}
		if atEOF {
			return 0, nil, io.EOF
		}
		return 0, nil, nil
	}
	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
