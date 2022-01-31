package fonction

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func ChooseWord1(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().Unix())
	nbrword := rand.Intn(len(strings.Split(string(content), "\n")) - 1)
	word := ""
	words := []string{}
	for _, chr := range content {
		if chr != 10 {
			word = word + string(chr)
		} else {
			words = append(words, word)
			word = ""
		}
	}
	return words[nbrword]
}
