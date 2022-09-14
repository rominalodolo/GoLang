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



type quizSpecs struct {
    Question string
    Correct string
    Answer2 string
	Answer3 string
	Answer4 string
}

type quizData struct {
	Q1 string
	Q2 string
	Q3 string
	Q4 string
	Q5 string
	Q6 string
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


		// create the variables 
    for _, line := range csvLines {
        quizS :=  quizSpecs {
            Question: line[0],
            Correct: line[1],
			Answer2: line[2],
			Answer3: line[3],
			Answer4: line[4],
        }
	fmt.Println("\n" + quizS.Question + ": \n" + quizS.Correct + " \n" + quizS.Answer2 + " \n" + quizS.Answer3 + " \n" + quizS.Answer4 + " \n")

	}
		csvColumns, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			fmt.Println(err)
		}  

		// run through the columns of the csv file till 
	for _, column := range csvColumns {
		quizD := quizData {
			Q1: column[0],
			Q2: column[1],
			Q3: column[2],
			Q4: column[3],
			Q5: column[4],
			Q6: column[5],
		}
		

		fmt.Println(quizD.Q1)
		
		
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
