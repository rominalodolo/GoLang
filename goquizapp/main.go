/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

// import "github.com/golang/goquizapp/cmd"
import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var tmpl *template.Template

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Welcome to the South African Whales and Dolphin Quiz!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

type CsvLine struct {
	Column1 string
}

func main() {
	// cmd.Execute()

	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	// filename := "quiz.csv"

	fmt.Println("You scored higher than 60% of all quizzers")

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
