package backpack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriority(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    rune
		expected int
	}{
		{'p', 16},
		{'L', 38},
		{'P', 42},
		{'v', 22},
		{'t', 20},
		{'s', 19},
	}

	for _, test := range tests {
		assert.Equal(Priority(test.input), test.expected)
	}
}

func TestFindDuplicates(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []string
		expected rune
	}{
		{[]string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, 'p'},
		{[]string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}, 'L'},
		{[]string{"PmmdzqPrV", "vPwwTWBwg"}, 'P'},
		{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}, 'r'},
		{[]string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}, 'Z'},
	}

	for _, test := range tests {
		assert.Equal(FindDuplicate(test.input...), test.expected)
	}
}
