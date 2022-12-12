package backpack

import (
	"sort"
	"strings"
)

func FindDuplicate(strs ...string) rune {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) >= len(strs[j])
	})
	totalStrings := len(strs)
	for _, c := range strs[0] {
		occurrences := 1
		for _, str := range strs[1:] {
			if strings.ContainsRune(str, c) {
				occurrences++
			}
		}
		if occurrences == totalStrings {
			return c
		}
	}

	panic("No duplicate found")
}

func Priority(c rune) int {
	if c > 'a' {
		return int(c-'a') + 1
	}

	return int(c-'A') + 27
}
