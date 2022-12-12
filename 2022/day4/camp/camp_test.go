package camp

import (
	"testing"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	input := "2-4"
	expected := Assignment{sections: *hashset.New(2, 3, 4)}

	actual := New(input)
	assert.Equal(actual.String(), expected.String())
}

func TestAreFullyOverlapping(t *testing.T) {
	assert := assert.New(t)

	assert.True(AreFullyOverlapping(New("2-6"), New("2-4")))
}
