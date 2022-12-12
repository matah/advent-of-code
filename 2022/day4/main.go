package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input
var input []byte

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}
