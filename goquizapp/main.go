/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	// "errors"
	"fmt"
	// "html/template"
	// "io"
	// "net/http"
	"os"
	"encoding/csv"
)



type quizData struct {
    Question string
    Correct string
    Answer2 string
	Answer3 string
	Answer4 string
}


func main() {

	csvFile, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Welcome to the Whales and Dolphin Quiz (South Africa EDITION) \n ")
	defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }    


    for _, line := range csvLines {
		if 
			quiz :=  quizData{
				Question: line[0],
				Correct: line[1],
				Answer2: line[2],
				Answer3: line[3],
				Answer4: line[4],
			}


        fmt.Println("\n" + quiz.Question + ": \n" + quiz.Correct + " \n" + quiz.Answer2 + " \n" + quiz.Answer3 + " \n" + quiz.Answer4 + " \n")
		

		
    }

    // var then variable name then variable type
    var first string
  
    // Taking input from user
    fmt.Scanln(&first)
    fmt.Println("Enter Second Last Name: ")
    var second string
    fmt.Scanln(&second)
  
    // Print function is used to
    // display output in the same line
    fmt.Print("Your Full Name is: ")
  
    // Addition of two string
    fmt.Print(first + " " + second)

	fmt.Println("\nYou scored higher than 60% of all quizzers")

}


