package io

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/matah/advent-of-code/day5/engine"
)

func ParseInitialState(input []byte) []engine.Stack {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	var (
		lines []string
	)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		lines = append(lines, line)
	}

	_, max := stackLimits(lines[len(lines)-1])

	stacks := make([]engine.Stack, max)
	for li := len(lines) - 2; li >= 0; li-- {
		for si := 0; si < max; si++ {
			idx := si*4 + 1
			currentLine := lines[li]
			if idx >= len(currentLine) {
				break
			}

			if unicode.IsLetter(rune(currentLine[idx])) {
				stacks[si].Append(rune(currentLine[idx]))
			}
		}
	}
	return stacks
}

func ParseMove(input string) (*engine.Move, error) {
	var moves, start, end int
	_, err := fmt.Sscanf(input, "move %d from %d to %d", &moves, &start, &end)
	if err != nil {
		return nil, err
	}

	return engine.NewMove(moves, start, end), nil
}

func stackLimits(input string) (int, int) {
	ids := strings.Split(strings.TrimSpace(input), " ")
	min, err := strconv.Atoi(ids[0])
	if err != nil {
		panic("Cannot parse beginning id for stack")
	}

	max, err := strconv.Atoi(ids[len(ids)-1])
	if err != nil {
		panic("Cannot parse ending id for stack")
	}

	return min, max
}
