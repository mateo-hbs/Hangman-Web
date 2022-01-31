package fonction

func Newletter(word, gameword, letter string) string {
	newgameword := []string{}
	check := true
	for _, k := range gameword {
		newgameword = append(newgameword, string(k))
	}
	for i := 0; i < len(newgameword); i++ {
		if letter == string(word[i]) {
			if letter != string(gameword[i]) {
				newgameword[i] = letter
			} else if check == true {
				check = false
			}
		}
	}
	gameword = ""
	for _, j := range newgameword {
		gameword = gameword + j
	}
	return gameword
}
