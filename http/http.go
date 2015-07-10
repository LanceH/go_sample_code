package main

import (
  "html/template"
	"math/rand"
	"net/http"
)

var pieces []string = []string{"rock", "paper", "scissors"}
var win, loss, draw int

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/rock", rock)
	http.HandleFunc("/paper", paper)
	http.HandleFunc("/scissors", scissors)

	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	result := "Choose Wisely"
	t, _ := template.ParseFiles("root.html")
	p := &Page{W: win, L: loss, D: draw, Result: result}
	t.Execute(w,p)
}

func rock(w http.ResponseWriter, r *http.Request) {
	computer := pieces[rand.Intn(3)]
	result := ""
	switch computer {
	case "rock":
		draw++
		result = "DRAW! Rock - Rock"
	case "paper":
		loss++
		result = "LOSE! Paper covers Rock"
	case "scissors":
		win++
		result = "WIN! Rock breaks Scissors"
	}
	t, _ := template.ParseFiles("root.html")
	p := &Page{W: win, L: loss, D: draw, Result: result}
	t.Execute(w,p)
}

func paper(w http.ResponseWriter, r *http.Request) {
	computer := pieces[rand.Intn(3)]
	result := ""
	switch computer {
	case "rock":
		win++
		result = "WIN! Paper covers Rock"
	case "paper":
		draw++
		result = "DRAW! Paper - Paper"
	case "scissors":
		loss++
		result = "LOSE! Scissors cut Paper"
	}
	t, _ := template.ParseFiles("root.html")
	p := &Page{W: win, L: loss, D: draw, Result: result}
	t.Execute(w,p)
}

func scissors(w http.ResponseWriter, r *http.Request) {
	computer := pieces[rand.Intn(3)]
	result := ""
	switch computer {
	case "rock":
		loss++
		result = "LOSE! Rock breaks scissors"
	case "paper":
		win++
		result = "WIN! Scissors cut paper"
	case "scissors":
		draw++
		result = "DRAW! Scissors - Scissors"
	}
	t, _ := template.ParseFiles("root.html")
	p := &Page{W: win, L: loss, D: draw, Result: result}
	t.Execute(w,p)
}

type Page struct {
	W int
	L int
	D int
	Result string
}
