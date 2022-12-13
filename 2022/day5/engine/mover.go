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

func (s *Stack) Pop() (rune, bool) {
	size := len(s.content)
	if size == 0 {
		return 0, false
	}

	elem := s.content[size-1]
	s.content = s.content[:size-1]
	return elem, true
}

func (s *Stack) Peek() rune {
	size := len(s.content)
	if size == 0 {
		return ' '
	}

	return s.content[size-1]
}

type Version int

const (
	Mover9000 Version = iota
	Mover9001
)

type Mover struct {
	version Version
	state   []Stack
}

func NewMover(initialState []Stack, ver Version) *Mover {
	return &Mover{
		state:   initialState,
		version: ver,
	}
}

func (m *Mover) Execute(move *Move) {
	switch m.version {
	case Mover9000:
		m.mover9000Execute(move)
	case Mover9001:
		m.mover9001Execute(move)
	default:
		panic("Unexpected version")
	}
}

func (m *Mover) mover9001Execute(move *Move) {
	var elems []rune
	for i := 0; i < move.numberOfMoves; i++ {
		elem, success := m.state[move.start-1].Pop()
		if !success {
			break
		}

		elems = append(elems, elem)
	}

	for i := len(elems) - 1; i >= 0; i-- {
		m.state[move.end-1].Append(elems[i])
	}
}

func (m *Mover) mover9000Execute(move *Move) {
	for i := 0; i < move.numberOfMoves; i++ {
		elem, success := m.state[move.start-1].Pop()
		if !success {
			break
		}
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
