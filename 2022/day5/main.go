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
	var (
		part1Total int
		part2Total int
	)

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}
