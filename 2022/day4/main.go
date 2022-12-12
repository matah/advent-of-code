package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	"github.com/matah/advent-of-code/day4/camp"
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

		pairs := strings.Split(line, ",")
		if camp.AreFullyOverlapping(camp.New(pairs[0]), camp.New(pairs[1])) {
			part1Total++
		}

		if camp.AreOverlapping(camp.New(pairs[0]), camp.New(pairs[1])) {
			part2Total++
		}
	}

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}
