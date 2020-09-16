package main

import "flag"

func main() {

	//create flag in CLI that prompts for a string representing a name of a csv file
	//helper text asks for a csv in our chosen format
	csvFilename := flag.String("csv", "problems.csv", "a csv in the format of 'question,answer'")

	flag.Parse()

	
	_ = csvFilename
}