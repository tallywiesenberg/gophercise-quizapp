package main

import ( 
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	//create flag in CLI that prompts for a string representing a name of a csv file
	//helper text asks for a csv in our chosen format
	csvFilename := flag.String("csv", "questions.csv", "a csv in the format of 'question,answer'")
	flag.Parse()

	
	file, err := os.Open(*csvFilename)
	if err != nil {
		//return a formatted string for client troubleshooting
		exit(fmt.Sprintf("Failed to open the csv file: %s \n", *csvFilename))
		os.Exit(1)
	}

	r :=csv.NewReader(file)
	lines, err :=r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	fmt.Println(lines)

}

func exit(msg string) {
	fmt.Printf(msg) //troubleshooting for the user
	os.Exit(1)
}