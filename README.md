# GoLang

Go handbook [Download PDF](https://thevalleyofcode.pages.dev/go-handbook.pdf)

GoLang Courses: 
- [Free Code Camp](https://www.youtube.com/watch?v=jFfo23yIWac&t=19s)
- [Tech World with Nana](https://www.youtube.com/watch?v=yyUHQIec83I&t=279s)
- [Golang Quiz App...](https://www.youtube.com/watch?v=TVHGxz6tn2M&t=960s)
- [GoLang Simple TODO App using HTML Templates](https://www.youtube.com/watch?v=UeUDGEa0AX4)


> Developed at Google and made open source in 2009 
> 
> It's a multithread language 
> 
> Designed to run on multiple cores and built to support concurrency
> 
> Concurrency in Go is cheap and easy
> 
> for performant applications
> 
> Docker, Vault, Cockroach is written in Go 


GoLang projects: Hello world starter 

### Hello World Starter App

![hellowithgo](https://user-images.githubusercontent.com/83961643/188650720-72aad107-2726-47d5-ad9b-61e73d14dc2a.jpeg)

![gohelp](https://user-images.githubusercontent.com/83961643/188650736-feb1857f-7bbf-472f-bb02-507e0f184234.jpeg)




### Go Quiz App 

The problem: 
- The task is to build a super simple quiz with a few questions and a few alternatives for each question. With one correct answer per question.
- Backend - Golang
- Database - Just in-memory , so no database
- REST API or gRPC
- CLI that talks with the API, preferably using https://github.com/spf13/cobra ( as cli framework )
- [cobra-cli Readme](https://github.com/spf13/cobra-cli/blob/main/README.md)

Installing Cobra 

![cobra-cli](https://user-images.githubusercontent.com/83961643/189537596-b82cd6e1-dbe5-4425-ac40-e131b34c8c73.jpeg)
![cobrainstall](https://user-images.githubusercontent.com/83961643/189361475-3538d69b-9c9b-4aee-8bd9-c32c602ec479.jpeg)

Facing issues with init but go it to work and get the Go project to initialise and run
![running](https://user-images.githubusercontent.com/83961643/189537612-7e245eba-8e56-43b9-be24-6db05023c1a0.jpeg)


Testing the application and running "Hello World!" 
![helloworldtest](https://user-images.githubusercontent.com/83961643/189537623-ea3a95b3-06f7-47b1-ba0d-a9329a9eba1a.jpeg)


[fmt package](https://pkg.go.dev/fmt@go1.19.1#hdr-Printing)

**Package fmt** implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.


[io package](https://pkg.go.dev/io)

**Package io** provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package os, into shared public interfaces that abstract the functionality, plus some other related primitives.

Because these interfaces and primitives wrap lower-level operations with various implementations, unless otherwise informed clients should not assume they are safe for parallel execution.



Being able to print in the terminal 

![print](https://user-images.githubusercontent.com/83961643/189546280-eb36542f-3ca2-408c-99f1-0561e694f203.jpeg)


[Create HTTP Server](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go) | http://localhost:3333

![http](https://user-images.githubusercontent.com/83961643/189546279-f876921a-bdf6-4ae2-9058-9e98006a4ba1.jpeg)



[template package](https://pkg.go.dev/html/template)

**Package template** (html/template) implements data-driven templates for generating HTML output safe against code injection. It provides the same interface as package text/template and should be used instead of text/template whenever the output is HTML.

![controlstruc](https://user-images.githubusercontent.com/83961643/189873380-cf0a995e-76e9-4df1-a478-ee374d91c4ab.jpeg)
[reference: Go Web Examples](https://gowebexamples.com/templates/)



Decoding a csv file in [go](https://pkg.go.dev/encoding/csv)

Trying to run the csv file but there's errors in how I put together the csv structure and questions. 
![runningthecsvfile](https://user-images.githubusercontent.com/83961643/189903221-7515d9b0-0d5e-4a82-b2c7-1205a4019da4.jpeg)
[code reference](https://webdamn.com/how-to-read-csv-file-using-golang/)


Running through the csv file: (Working csv values)

![readingcsv](https://user-images.githubusercontent.com/83961643/189907659-4f6e28ac-c2ef-440d-b85b-c3c2b197cf35.jpeg)


**The Problem**

_User stories/Use cases_

* _User should be presented questions with a number of answers_

* _User should be able to select just one answer per question_

* _User should be able to answer all the questions and then post his/her answers and get back how many correct answers there had and be displayed to the user._

* _User should see how good he/she rated compared to others that have taken the quiz, "You scored higher than 60% of all quizzers"_



> I created a multiple choice quiz that has options between A->D (4 choices to choose from)
> 
> Only one right answer per question.
> 
> There are 6 quesions the program loops through but the user only sees 3 questions at a time.
> 
> When the quiz is completed the user will see their results with the answer. 
>
> And if the user has scored higher than 60%, the user will see the printed "You scored higher than 60% of all quizzers" 
> 
> The user will be asked if they would like to retry the quiz or if they would like to exit the application.



Last Edit: September 2022
