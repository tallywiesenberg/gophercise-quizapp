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

	//reading and parsing the csv into an array
	r :=csv.NewReader(file)
	lines, err :=r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	fmt.Println(problems)
	
	//COMMAND LINE OUTPUT AND INPUT w/ scanner
	for i, p := range problems {

		//output
		fmt.Printf("Problem #%d: %s = \n", 
		i+1, //question number starts at 1, not 0
		p.q)  //the question included

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("correct!")
		}
	}
}

func parseLines(lines [][]string) []problem {
	//each []problem slice is a row from the csv file
	//representing each row as a "problem" struct
	ret := make([]problem, len(lines))
	
	for i, line := range(lines) {
		ret[i] = problem {
			q: line[0],
			a: line[1],
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}


func exit(msg string) {
	fmt.Printf(msg) //troubleshooting for the user
	os.Exit(1)
}