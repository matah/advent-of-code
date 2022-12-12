package camp

import (
	"strconv"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

type Assignment struct {
	sections hashset.Set
}

func AreFullyOverlapping(assignment *Assignment, another *Assignment) bool {
	return assignment.sections.Contains(another.sections.Values()...) || another.sections.Contains(assignment.sections.Values()...)
}

func AreOverlapping(assignment *Assignment, another *Assignment) bool {
	return !assignment.sections.Intersection(&another.sections).Empty()
}

func New(str string) *Assignment {
	limits := strings.Split(str, "-")
	min, err := strconv.Atoi(limits[0])
	if err != nil {
		panic("incorrect input format")
	}

	max, err := strconv.Atoi(limits[1])
	if err != nil {
		panic("incorrect input format")
	}

	sections := hashset.New()
	for i := min; i <= max; i++ {
		sections.Add(i)
	}

	return &Assignment{
		sections: *sections,
	}
}

func (assignment *Assignment) String() string {
	return assignment.sections.String()
}
