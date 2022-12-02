package main

import (
	"bufio"
	"bytes"
	"day1/day1"
	_ "embed"
	"fmt"
	"log"
)

//go:embed input
var input []byte

func main() {
	entourage := day1.ElvenEntourage{
		Elves: []day1.Elf{},
	}

	id := uint64(0)
	e := day1.Elf{
		Identity:      id,
		TotalCalories: 0,
	}
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			e.AddSnack(line)
		} else {
			entourage.AddElf(e)
			id += 1
			e = day1.Elf{
				Identity:      id,
				TotalCalories: 0,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(entourage.GetTop(1))
	fmt.Println(entourage.GetTopSum(3))
}
