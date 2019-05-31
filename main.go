package main

import (
	"encoding/csv"
	"fmt"
	"os"
	// "log"
	// "io"
	"flag"
	"time"
)


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type question struct{
	question string
	answer string
}

func main() {
	fileName:= flag.String("fileName","problems","Enter file name")
	*fileName=*fileName+".csv"
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		// log.Fatal(err)
		exit(fmt.Sprintf("Failed to open csv file: %s",*fileName))
	}
	defer file.Close()

	r,err := csv.NewReader(file).ReadAll()
	if err!=nil{
		exit(fmt.Sprintf("Failed to parse csv file"))
	}

	questions:=getQuestions(r)
	correctCount:=askQuestions(questions)
	
	fmt.Printf("You have answered %v questions correctly.",correctCount)
}

func getQuestions(data[][]string) []question{
	q:=make([]question,len(data))
	for i,v:=range data {
		q[i]=question{question:v[0],answer:v[1]}
	}
	return q
}

func askQuestions(questions []question) int{
	var correctCount int
	an := make(chan string) 
	timer:=time.NewTimer(30*time.Second)
	
	questionsloop:for i,v:= range questions{
	go func(){
	var a string
	fmt.Printf("Question No %v: %v is ",i+1,v.question)
	fmt.Scan(&a)
	an<-a
	}()

	select{
	case <- timer.C:
		fmt.Println()
			break questionsloop;
	case answer:=<-an:
		if answer==v.answer{
			correctCount++
		}
	}
	} 
	
	return correctCount
}