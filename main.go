package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Starting execution...")

	inputFileName := flag.String("inFile", "input.csv", "a CSV of API requests to parse")
	outputFileName := flag.String("outFile", "output.csv", "the name of the file to write the output to")
	flag.Parse()

	fmt.Println("Reading File ", *inputFileName)
	fmt.Println("Preparing to write to file ", *outputFileName)

	//Check if the input file exists

	//Read one line

	//Parse out just the Method and path (case insensitive)

	//store in Key-value store with value 1 or increment existing occurrence

}
