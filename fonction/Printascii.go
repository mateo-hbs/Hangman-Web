package fonction

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func PrintASCII(word string) {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	letter := strings.Split(string(content), "\r\n")
	tabword := []byte{}
	for _, chr := range word {
		tabword = append(tabword, byte(chr))
	}
	for j := 0; j < 8; j++ {
		for i := 0; i < len(tabword); i += 1 {
			fmt.Print(letter[((int(tabword[i])-97)*9 + 586 + j)])
			fmt.Print(" ")
			time.Sleep(10 * time.Millisecond)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
