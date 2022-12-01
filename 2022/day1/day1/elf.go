package day1

import "strconv"

type Elf struct {
	Identity      uint64
	TotalCalories uint64
}

func (e *Elf) AddSnack(snack string) {
	v, err := strconv.ParseUint(snack, 10, 64)
	if err != nil {
		panic(err)
	}

	e.TotalCalories += v
}
