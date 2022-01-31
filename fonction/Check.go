package fonction

func Check(letter string, word string) bool {
	for i := 0; i < len(word); i++ {
		if letter == string(word[i]) {
			return true
		}
	}
	return false
}
