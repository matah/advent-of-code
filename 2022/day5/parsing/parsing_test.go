package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriority(t *testing.T) {
	assert := assert.New(t)

	input := `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

 move 1 from 2 to 1
`
	expected := []Stack{
		{content: []rune{'Z', 'N'}},
		{content: []rune{'M', 'C', 'D'}},
		{content: []rune{'P'}},
	}
	actual := ParseInitialState([]byte(input))
	assert.Equal(expected, actual)
}
