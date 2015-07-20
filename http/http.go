package main

import (
	"fmt"
  "html/template"
	"math/rand"
	"net/http"
	"time"
)

var pieces []string = []string{"rock", "paper", "scissors"}
var win, loss, draw int

func main() {
	rand.Seed( time.Now().UTC().UnixNano())
	http.HandleFunc("/", root)
	http.HandleFunc("/rock", rock)
	http.HandleFunc("/paper", paper)
	http.HandleFunc("/scissors", scissors)

	http.HandleFunc("/cookieset", cSet)
	http.HandleFunc("/cookieread", cRead)

	http.ListenAndServe(":8080", nil)
}

func cSet(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{Name: "test", Value: "Cookies!"}
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "cookie set")
}

func cRead(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("test")
	if err != nil {
		fmt.Fprintf(w, "No cookies")
	} else {
		fmt.Fprintf(w, "cookie: %s", cookie.Value)
	}
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

type User struct {
	id string
	Record *Record
}

type Record struct {
	W int
	L int
	D int
}

type Page struct {
	W int
	L int
	D int
	Result string
}
