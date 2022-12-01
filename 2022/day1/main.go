package main

import (
	"day1/day1"
	"fmt"
	"os"
)

func main() {

	const filename = "day1.input"
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	groups := day1.CalorieGroups(string(bytes))
	fmt.Println(day1.MaxCalories(groups))
	maxN := day1.MaxN(groups, 3)
	fmt.Println(sum(maxN...))
}

func sum(numbs ...int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}
