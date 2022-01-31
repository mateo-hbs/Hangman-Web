package fonction

import (
)

func SuggestedWord(letter string, word string) int {
	letter = ToUpper(letter)
	x := 2
	if letter != word {
		x = 0
	}
	if letter == word {
		x = 1
	}
	return x
}
