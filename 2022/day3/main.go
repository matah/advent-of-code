package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"

	"github.com/matah/advent-of-code/day3/backpack"
)

//go:embed input
var input []byte

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var (
		part1Prio, part2Prio, idx int
		backpacks                 [3]string
	)

	for scanner.Scan() {
		line := scanner.Text()

		half := len(line) / 2
		firstCompartment := line[:half]
		secondCompartment := line[half:]
		duplicate := backpack.FindDuplicate(firstCompartment, secondCompartment)
		part1Prio += backpack.Priority(duplicate)

		backpacks[idx%3] = line
		if idx%3 == 2 {
			// reset stuff
			part2Prio += backpack.Priority(backpack.FindDuplicate(backpacks[:]...))
		}
		idx++
	}

	fmt.Println(part1Prio)
	fmt.Println(part2Prio)
}
