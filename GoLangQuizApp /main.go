package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Quiz struct {
	Question string
	Answered bool
}

type PageData struct {
	Title string
	Quiz  []Quiz
}

func quiz(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title: "GoLang Quiz",
		Quiz: []Quiz{
			{Question: "What do whales breathe out of?", Answered: true},
			{Question: "How do whales regulate/create heat?", Answered: false},
			{Question: "Dolphins are very _________ animals and hunt in pods.", Answered: false},
			{Question: "A blue whale is the largest mammal in the ocean.", Answered: false},
			{Question: "A Bryde’s Whales’ diet consists of _____ .", Answered: false},
			{Question: "Humpback whales swim to which coast to have their babies?", Answered: true},
			{Question: "Bottle-nose dolphins are known for _______ .", Answered: true},
			{Question: "WHich dolphin spins in the air?", Answered: true},
			{Question: "What do whales breathe out of?", Answered: true},
		},
	}

	tmpl.Execute(w, data)

}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/quiz", quiz)

	log.Fatal(http.ListenAndServe(":9091", mux))
}