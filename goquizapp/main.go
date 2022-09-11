/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

// import "github.com/golang/goquizapp/cmd" 
import "fmt"
import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {
	// cmd.Execute()

	
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fmt.Println("You scored higher than 60% of all quizzers")

	log.Fatal(http.ListenAndServe(":9091", mux))

}
