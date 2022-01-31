package main

import (
	"hangman/fonction"
	"html/template"
	"net/http"
)

type Hangman struct {
	Valeur      string
	Name        string
	Level       string
	Life        int
	Gameword    string
	Hideword    string
	Lettersused []string
	Deadman     string
}

// url http://localhost:8080/

var templates = template.Must(template.ParseFiles("pages/index.html"))
var game = Hangman{Valeur: "Hangman Web", Name: "Player", Level: "", Life: 11, Gameword: "", Hideword: "", Lettersused: []string{}, Deadman: "static/pictures/deadman/happy.png"}

func homeHandler(w http.ResponseWriter, r *http.Request) { //function who associate a html file to a web page
	replay()
	a := template.Must(template.ParseFiles("pages/index.html"))
	err := a.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RulesFile(w http.ResponseWriter, r *http.Request) { //function who associate a html file to a web page
	a := template.Must(template.ParseFiles("pages/rules.html"))
	err := a.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutUs(w http.ResponseWriter, r *http.Request) { //function who associate a html file to a web page
	a := template.Must(template.ParseFiles("pages/credit.html"))
	err := a.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LevelDecider(w http.ResponseWriter, r *http.Request) { //function who associate a html file to a web page
	replay()
	game.Name = r.FormValue("name")
	if game.Name == "" {
		game.Name = "Player"
	}
	game.Level = r.FormValue("Level")
	if game.Level == ""{
		game.Level ="1"
	}
	a := template.Must(template.ParseFiles("pages/ChooseLevel.html"))
	err := a.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GamePage(w http.ResponseWriter, r *http.Request) {
	if game.Level == "1" { //if the user choose the level 1, the words will be choose in the words.txt file
		game.Hideword = fonction.ChooseWord1("words.txt")
		game.Hideword = fonction.ToUpper(game.Hideword)
		game.Gameword = fonction.Reveal_n_letters(game.Hideword)
		game.Gameword = fonction.ToUpper(game.Gameword)
	} else if game.Level == "2" {//if the user choose the level 2, the words will be choose in the words2.txt file
		game.Hideword = fonction.ChooseWord1("words2.txt")
		game.Hideword = fonction.ToUpper(game.Hideword)
		game.Gameword = fonction.Reveal_n_letters(game.Hideword)
		game.Gameword = fonction.ToUpper(game.Gameword)
	} else if game.Level == "3" { //if the user choose the level 3, the words will be choose in the words3.txt file
		game.Hideword = fonction.ChooseWord1("Words3.txt")
		game.Hideword = fonction.ToUpper(game.Hideword)
		game.Gameword = fonction.Reveal_n_letters(game.Hideword)
		game.Gameword = fonction.ToUpper(game.Gameword)
	}
	game.Level = ""
	letter := r.FormValue("letter")
	letter = fonction.ToUpper(letter)
	game.Lettersused = fonction.Letterused(game.Lettersused, letter)
	if len(letter) > 1 {
		if fonction.SuggestedWord(letter, game.Hideword) == 1{
			game.Gameword = game.Hideword
		}
	}
	if !fonction.Check(letter, game.Hideword) && game.Gameword != game.Hideword {
		if game.Life > 0 {
			game.Life = game.Life - 1
		}
		if game.Life < 10 { //print a picture for each life's status
			if game.Life == 9 {
				game.Deadman = "static/pictures/deadman/9HP.png"
			}
			if game.Life == 8 {
				game.Deadman = "static/pictures/deadman/8HP.png"
			}
			if game.Life == 7 {
				game.Deadman = "static/pictures/deadman/7HP.png"
			}
			if game.Life == 6 {
				game.Deadman = "static/pictures/deadman/6HP.png"
			}
			if game.Life == 5 {
				game.Deadman = "static/pictures/deadman/5HP.png"
			}
			if game.Life == 4 {
				game.Deadman = "static/pictures/deadman/4HP.png"
			}
			if game.Life == 3 {
				game.Deadman = "static/pictures/deadman/3HP.png"
			}
			if game.Life == 2 {
				game.Deadman = "static/pictures/deadman/2HP.png"
			}
			if game.Life == 1 {
				game.Deadman = "static/pictures/deadman/1HP.png"
			}
			if game.Life == 0 {
				game.Deadman = "static/pictures/deadman/0HP.png"
			}
		}
	}
	if game.Life == 0 {
		game.Gameword = "Mince tu as perdu le mot était : " + game.Hideword
	} else {
		game.Gameword = fonction.Newletter(game.Hideword, game.Gameword, letter)
	}
	letter = ""
	if game.Gameword == game.Hideword {
		game.Gameword = "Bravo tu as gagné"
	}
	a := template.Must(template.ParseFiles("pages/game.html"))
	err := a.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func replay() { //this function as in goal to reset all the params when you leave de Game Page
	game.Name = "player"
	game.Level = ""
	game.Life = 11
	game.Gameword = ""
	game.Hideword = ""
	game.Lettersused = []string{}
	game.Deadman = ""
}

func main() {
	http.HandleFunc("/", homeHandler) //each HandleFunc are associated with a html file to open 
	http.HandleFunc("/ChooseLevel", LevelDecider) // the right page on web server 
	http.HandleFunc("/Rules", RulesFile)
	http.HandleFunc("/Credit", AboutUs)
	http.HandleFunc("/Game", GamePage)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./pages"))))
	http.ListenAndServe(":8080", nil) 	//defind the url of the serv :
										// url http://localhost:8080/
}
