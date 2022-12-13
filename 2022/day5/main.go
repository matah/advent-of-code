package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"

	"github.com/matah/advent-of-code/day5/engine"
	"github.com/matah/advent-of-code/day5/io"
)

//go:embed input
var input []byte

func main() {
	stacks := io.ParseInitialState(input)
	mover := engine.NewMover(stacks)

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		move, err := io.ParseMove(line)
		if err != nil {
			continue
		}

		mover.Execute(move)
	}

	fmt.Println(mover.State())
}
