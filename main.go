package main

import (
	"encoding/csv"
	"fmt"
	"os"
	// "log"
	// "io"
	"flag"
)


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
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

	// fmt.Println(r)
	fmt.Printf("%T\n",r)
	var correctCount int
	for i,v:=range r{
		var a string
	// fmt.Printf("This is the question %s and this is the answer %s \n",v[0],v[1])	
	fmt.Printf("Question No %v: %v is ",i,v[0])
	fmt.Scan(&a)
	if(a==v[1]){
		correctCount++
	}
	}
	fmt.Printf("You have answered %v questions correctly.",correctCount)
}
