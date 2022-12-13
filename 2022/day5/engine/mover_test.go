package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	assert := assert.New(t)

	stacks := []Stack{
		*NewStack([]rune{'Z', 'N'}),
		*NewStack([]rune{'M', 'C', 'D'}),
		*NewStack([]rune{'P'}),
	}

	target := NewMover(stacks)
	target.Execute(NewMove(1, 2, 1))
	state := target.State()
	assert.Equal("DCP", state)

	target.Execute(NewMove(3, 1, 3))
	state = target.State()
	assert.Equal(" CZ", state)

	target.Execute(NewMove(2, 2, 1))
	state = target.State()
	assert.Equal("M Z", state)

	target.Execute(NewMove(1, 1, 2))
	state = target.State()
	assert.Equal("CMZ", state)
}
