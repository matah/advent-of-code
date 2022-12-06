package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGestureParsing_GivenValidInput(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    byte
		expected Gesture
	}{
		{'A', Rock},
		{'X', Rock},
		{'B', Paper},
		{'Y', Paper},
		{'C', Scissors},
		{'Z', Scissors},
	}

	for _, test := range tests {
		assert.Equal(ToGesture(test.input), test.expected)
	}
}

func TestPlayRound_DrawScenarios(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		opponent Gesture
		me       Gesture
		expected Result
	}{
		{Rock, Rock, Draw},
		{Paper, Paper, Draw},
		{Scissors, Scissors, Draw},
	}

	for _, test := range tests {
		assert.Equal(PlayRound(test.opponent, test.me), test.expected)
	}
}

func TestPlayRound_WinScenarios(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		opponent Gesture
		me       Gesture
		expected Result
	}{
		{Rock, Paper, Win},
		{Paper, Scissors, Win},
		{Scissors, Rock, Win},
	}

	for _, test := range tests {
		assert.Equal(PlayRound(test.opponent, test.me), test.expected)
	}
}

func TestPlayRound_LossScenarios(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		opponent Gesture
		me       Gesture
		expected Result
	}{
		{Rock, Scissors, Loss},
		{Paper, Rock, Loss},
		{Scissors, Paper, Loss},
	}

	for _, test := range tests {
		assert.Equal(PlayRound(test.opponent, test.me), test.expected)
	}
}

func TestGetResult(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    byte
		expected Result
	}{
		{'X', Loss},
		{'Y', Draw},
		{'Z', Win},
	}

	for _, test := range tests {
		assert.Equal(ToResult(test.input), test.expected)
	}
}

func TestGetMyMove_DrawScenarios(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		opponent Gesture
		result   Result
		expected Gesture
	}{
		{Rock, Draw, Rock},
		{Paper, Draw, Paper},
		{Scissors, Draw, Scissors},
	}

	for _, test := range tests {
		assert.Equal(GetMyMove(test.opponent, test.result), test.expected)
	}
}

func TestGetMyMove_WinScenarios(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		opponent Gesture
		result   Result
		expected Gesture
	}{
		{Rock, Win, Paper},
		{Paper, Win, Scissors},
		{Scissors, Win, Rock},
	}

	for _, test := range tests {
		assert.Equal(GetMyMove(test.opponent, test.result), test.expected)
	}
}

func TestGetMyMove_LossScenarios(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		opponent Gesture
		result   Result
		expected Gesture
	}{
		{Rock, Loss, Scissors},
		{Paper, Loss, Rock},
		{Scissors, Loss, Paper},
	}

	for _, test := range tests {
		assert.Equal(GetMyMove(test.opponent, test.result), test.expected)
	}
}
