package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input []byte

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		half := len(line) / 2
		firstCompartment := line[:half]
		secondCompartment := line[half:]
		duplicate := findDuplicate(firstCompartment, secondCompartment)
		total += priority(duplicate)
	}

	fmt.Println(total)
}

func findDuplicate(s1 string, s2 string) rune {
	for _, c := range s1 {
		if strings.ContainsRune(s2, c) {
			return c
		}
	}

	panic("No duplicate found")
}

func priority(c rune) int {
	if c > 'a' {
		return int(c-'a') + 1
	}

	return int(c-'A') + 27
}
