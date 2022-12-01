package day1

import (
	"strconv"
	"strings"
)

func CalorieGroups(input string) [][]int {
	lines := strings.Split(input, "\n")
	var groups [][]int
	var group []int

	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = nil
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			group = append(group, value)
		}
	}

	return groups
}

func caloriesForGroup(group []int) int {
	result := 0
	for _, v := range group {
		result += v
	}
	return result
}

func MaxCalories(groups [][]int) int {
	max := 0
	for _, group := range groups {
		if groupCalories := caloriesForGroup(group); groupCalories > max {
			max = groupCalories
		}
	}
	return max
}
