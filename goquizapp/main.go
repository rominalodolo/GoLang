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
	"encoding/csv"
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

type quizData struct {
    Question string
    Correct string
    Answer2 string
	Answer3 string
	Answer4 string
}

func main() {
	// cmd.Execute()

	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	// quizData := "quiz.csv"

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


	csvFile, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }    
    for _, line := range csvLines {
        quiz :=  quizData{
            Question: line[0],
            Age: line[1],
			City: line[2],
        }
        fmt.Println(emp.Name + " " + emp.Age + " " + emp.City)
    }

}
