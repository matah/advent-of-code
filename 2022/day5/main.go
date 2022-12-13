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
	initState9000 := make([]engine.Stack, len(stacks))
	initState9001 := make([]engine.Stack, len(stacks))

	copy(initState9000, stacks)
	copy(initState9001, stacks)

	mover9000 := engine.NewMover(initState9000, engine.Mover9000)
	mover9001 := engine.NewMover(initState9001, engine.Mover9001)

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		move, err := io.ParseMove(line)
		if err != nil {
			continue
		}

		mover9000.Execute(move)
		mover9001.Execute(move)
	}

	fmt.Println(mover9000.State())
	fmt.Println(mover9001.State())
}
