package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGestureScore(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    Gesture
		expected int
	}{
		{Rock, 1},
		{Paper, 2},
		{Scissors, 3},
	}

	for _, test := range tests {
		assert.Equal(test.input.Score(), test.expected)
	}
}

func TestResultScore(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    Result
		expected int
	}{
		{Win, 6},
		{Loss, 0},
		{Draw, 3},
	}

	for _, test := range tests {
		assert.Equal(test.input.Score(), test.expected)
	}
}
