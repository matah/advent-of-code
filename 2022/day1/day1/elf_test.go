package day1

import "testing"

func TestAddSnack(t *testing.T) {
	e := Elf{TotalCalories: 0, Identity: 0}
	input := "1234"
	expected := 1234

	e.AddSnack(input)

	if e.TotalCalories != 1234 {
		t.Fatalf("AddSnack() = %v, want %v", e.TotalCalories, expected)
	}
}
