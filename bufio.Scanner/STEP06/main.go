package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func main() {
	input := "abcdefghijkl"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		return 0, nil, errors.New("bad luck")
	}
	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Printf("error: %s\n", scanner.Err())
	}
}
