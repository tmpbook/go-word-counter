package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	input := "foo|bar"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 防止无限循环的方法
		// if atEOF && len(data) == 0 {
		// 	return 0, nil, nil
		// }
		if i := bytes.IndexByte(data, '|'); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data[:len(data)], nil
		}
		return 0, nil, nil
	}
	scanner.Split(split)
	for scanner.Scan() {
		if scanner.Text() != "" {
			fmt.Println(scanner.Text())
		}
	}
}
