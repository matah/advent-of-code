package day1

import (
	"testing"
)

func TestAddElf(t *testing.T) {
	e1 := Elf{TotalCalories: 123, Identity: 0}
	e2 := Elf{TotalCalories: 123, Identity: 0}

	target := ElvenEntourage{Elves: []Elf{}}

	target.AddElf(e1)
	target.AddElf(e2)

	if len(target.Elves) != 2 {
		t.Fatalf("AddSnack() = %v, want %v", len(target.Elves), 2)
	}
}
