package engine

func ToGesture(b byte) Gesture {
	switch b {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	default:
		panic("Unknown input")
	}
}

func PlayRound(opponent Gesture, me Gesture) Result {
	if opponent == me {
		return Draw
	}

	if opponent == Rock && me == Scissors ||
		opponent == Paper && me == Rock ||
		opponent == Scissors && me == Paper {
		return Loss
	}

	return Win
}

func ToResult(b byte) Result {
	switch b {
	case 'X':
		return Loss
	case 'Y':
		return Draw
	case 'Z':
		return Win
	default:
		panic("Unknown input")
	}
}

func GetMyMove(opponent Gesture, result Result) Gesture {
	if result == Draw {
		return opponent
	}

	if result == Win {
		return opponent.Opposite()
	}

	return opponent.Opposite().Opposite()
}
