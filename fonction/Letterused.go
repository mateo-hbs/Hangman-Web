package fonction

func Letterused(letteruse []string, letter string) []string {
	check := false
	for _, i := range letteruse {
		if i == letter {
			check = true
		}
	}
	if !check {
		letteruse = append(letteruse, letter)
	}
	return letteruse
}
