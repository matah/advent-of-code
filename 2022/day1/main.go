package main

import (
	"bufio"
	"day1/day1"
	"fmt"
	"log"
	"os"
)

func main() {
	const filename = "day1.input"
	entourage := day1.ElvenEntourage{
		Elves: []day1.Elf{},
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	id := uint64(0)
	e := day1.Elf{
		Identity:      id,
		TotalCalories: 0,
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if l := scanner.Text(); l != "" {
			e.AddSnack(l)
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
