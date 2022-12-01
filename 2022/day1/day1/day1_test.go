package day1

import "testing"

func TestHelloWorld(t *testing.T) {
	input := "6146\n5646\n\n30144\n31365\n"
	expected := [][]int{{6146, 5646}, {30144, 31365}}
	if observed := CalorieGroups(input); !testEq(observed, expected) {
		t.Fatalf("CalorieGroups() = %v, want %v", observed, expected)
	}
}

func testEq(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, inner := range a {
		for j, value := range inner {
			if value != b[i][j] {
				return false
			}
		}

	}
	return true
}

func BenchmarkCalorieGroups(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	input := "6146\n5646\n\n30144\n31365\n"
	for i := 0; i < b.N; i++ {
		CalorieGroups(input)
	}
}
