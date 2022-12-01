package day1

import (
	"sort"
)

type ElvenEntourage struct {
	Elves []Elf
}

func (entourage *ElvenEntourage) AddElf(e Elf) {
	entourage.Elves = append(entourage.Elves, e)
}

func (entourage *ElvenEntourage) GetTop(n uint64) []Elf {
	elves := make([]Elf, len(entourage.Elves))
	copy(elves, entourage.Elves)

	sort.Slice(elves, func(i, j int) bool {
		return entourage.Elves[i].TotalCalories > entourage.Elves[j].TotalCalories
	})

	return elves[:n]
}

func (entourage *ElvenEntourage) GetTopSum(n uint64) uint64 {
	sorted := entourage.GetTop(n)
	result := uint64(0)
	for _, v := range sorted {
		result += v.TotalCalories
	}
	return result
}
