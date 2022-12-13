package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute9000(t *testing.T) {
	assert := assert.New(t)

	stacks := []Stack{
		*NewStack([]rune{'Z', 'N'}),
		*NewStack([]rune{'M', 'C', 'D'}),
		*NewStack([]rune{'P'}),
	}

	target := NewMover(stacks, Mover9000)
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

func TestExecute9001(t *testing.T) {
	assert := assert.New(t)

	stacks := []Stack{
		*NewStack([]rune{'Z', 'N'}),
		*NewStack([]rune{'M', 'C', 'D'}),
		*NewStack([]rune{'P'}),
	}

	target := NewMover(stacks, Mover9001)
	target.Execute(NewMove(1, 2, 1))
	state := target.State()
	assert.Equal("DCP", state)

	target.Execute(NewMove(3, 1, 3))
	state = target.State()
	assert.Equal(" CD", state)

	target.Execute(NewMove(2, 2, 1))
	state = target.State()
	assert.Equal("C D", state)

	target.Execute(NewMove(1, 1, 2))
	state = target.State()
	assert.Equal("MCD", state)
}
