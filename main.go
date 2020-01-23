package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Starting execution...")

	inputFileName := flag.String("inFile", "input.csv", "a CSV of API requests to parse")
	outputFileName := flag.String("outFile", "output.csv", "the name of the file to write the output to")
	flag.Parse()

	fmt.Println("Reading File ", *inputFileName)
	fmt.Println("Preparing to write to file ", *outputFileName)

	//Check if the input file exists
	//inReader, err := os.Open(*inputFileName)

	// if err != nil {
	// 	panic(err)
	// }

	//Read one line

	//Parse out just the Method and path (case insensitive)

	//store in Key-value store with value 1 or increment existing occurrence

}

//ParseURL parses a thing
func ParseURL(url string) string {
	parts := strings.SplitAfter(url, "?")
	urlSansQuestionMark := strings.TrimSuffix(parts[0], "?")
	urlSansSlash := strings.TrimSuffix(urlSansQuestionMark, "/")
	return urlSansSlash
}
