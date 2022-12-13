package engine

type Stack struct {
	content []rune
}

func NewStack(values []rune) *Stack {
	return &Stack{content: values}
}

func (s *Stack) Append(value rune) {
	s.content = append(s.content, value)
}

func (s *Stack) Pop() rune {
	size := len(s.content)
	if size == 0 {
		panic("Cannot pop from empty stack")
	}

	elem := s.content[size-1]
	s.content = s.content[:size-1]
	return elem
}

func (s *Stack) Peek() rune {
	size := len(s.content)
	if size == 0 {
		return ' '
	}

	return s.content[size-1]
}

type Mover struct {
	state []Stack
}

func NewMover(initialState []Stack) *Mover {
	return &Mover{
		state: initialState,
	}
}

func (m *Mover) Execute(move *Move) {
	for i := 0; i < move.numberOfMoves; i++ {
		elem := m.state[move.start-1].Pop()
		m.state[move.end-1].Append(elem)
	}
}

func (m *Mover) State() string {
	var tops []rune
	for _, s := range m.state {
		tops = append(tops, s.Peek())
	}

	return string(tops)
}

type Move struct {
	numberOfMoves int
	start         int
	end           int
}

func NewMove(numberOfMoves int, start int, end int) *Move {
	return &Move{
		numberOfMoves: numberOfMoves,
		start:         start,
		end:           end,
	}
}
