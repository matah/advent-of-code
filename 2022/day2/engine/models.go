package engine

type Gesture int

const (
	Rock Gesture = iota
	Paper
	Scissors
)

func (g Gesture) String() string {
	return [...]string{"Rock", "Paper", "Scissors"}[g]
}

func (g Gesture) Opposite() Gesture {
	switch g {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		panic("Unknown gesture")
	}
}

func (g Gesture) Score() int {
	return [...]int{1, 2, 3}[g]
}

type Result int

const (
	Win Result = iota
	Loss
	Draw
)

func (r Result) String() string {
	return [...]string{"Win", "Loss", "Draw"}[r]
}

func (r Result) Score() int {
	return [...]int{6, 0, 3}[r]
}
