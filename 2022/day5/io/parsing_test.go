package io

import (
	"testing"

	"github.com/matah/advent-of-code/day5/engine"
	"github.com/stretchr/testify/assert"
)

func TestParseInitialState(t *testing.T) {
	assert := assert.New(t)

	input := `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3 

 move 1 from 2 to 1
`
	expected := []engine.Stack{
		*engine.NewStack([]rune{'Z', 'N'}),
		*engine.NewStack([]rune{'M', 'C', 'D'}),
		*engine.NewStack([]rune{'P'}),
	}
	actual := ParseInitialState([]byte(input))
	assert.Equal(expected, actual)
}

func TestParseMoveGivenCorrectInput(t *testing.T) {
	assert := assert.New(t)

	input := "move 1 from 2 to 1"
	expected := engine.NewMove(1, 2, 1)
	actual, err := ParseMove(input)
	assert.Equal(expected, actual)
	assert.Nil(err)
}

func TestParseMoveGivenWrongInput(t *testing.T) {
	assert := assert.New(t)

	input := "bla bla bla"
	move, err := ParseMove(input)
	assert.Nil(move)
	assert.NotNil(err)
}
