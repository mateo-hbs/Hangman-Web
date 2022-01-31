package fonction

func ToUpper(word string) string {
	tab := []rune(word)
	for i := 0; i < len(word); i++ {
		if tab[i] >= 'a' && tab[i] <= 'z' {
			tab[i] = tab[i] - 32
		}
	}
	return string(tab)
}
