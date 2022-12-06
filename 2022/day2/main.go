package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"

	"github.com/matah/advent-of-code/2022/day2/engine"
)

//go:embed input
var input []byte

func calculatePart1Score(line string) int {
	opponent := engine.ToGesture(line[0])
	me := engine.ToGesture(line[2])
	result := engine.PlayRound(opponent, me)

	return result.Score() + me.Score()
}

func calculatePart2Score(line string) int {
	opponent := engine.ToGesture(line[0])
	result := engine.ToResult(line[2])
	me := engine.GetMyMove(opponent, result)

	return result.Score() + me.Score()
}

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	part1Score := 0
	part2Score := 0
	for scanner.Scan() {
		line := scanner.Text()
		part1Score += calculatePart1Score(line)
		part2Score += calculatePart2Score(line)
	}

	fmt.Println(part1Score)
	fmt.Println(part2Score)
}
