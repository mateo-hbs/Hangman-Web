package fonction

import (
	"math/rand"
	"time"
)

func Reveal_n_letters(word string) string {
	var word_letters []string
	var mot_cache []string
	var lettres_trouvees []string
	var new_word string
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(word); i++ {
		word_letters = append(word_letters, string(word[i]))
	}
	for {
		if len(lettres_trouvees) >= len(word_letters)/2-1 {
			break
		}
		X := word_letters[rand.Intn(len(word))]
		lettres_trouvees = append(lettres_trouvees, X)
	}
	for i := 0; i < len(word_letters); i++ {
		var found bool = false
		for j := 0; j < len(lettres_trouvees); j++ {
			if word_letters[i] == lettres_trouvees[j] {
				found = true
				break
			}
		}
		if found {
			mot_cache = append(mot_cache, word_letters[i])
		} else {
			mot_cache = append(mot_cache, "_")
		}
	}
	for _, i := range mot_cache {
		new_word = new_word + i
	}
	return new_word

}
