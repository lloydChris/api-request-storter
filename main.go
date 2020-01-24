package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
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
	fileReader, err := os.Open(*inputFileName)
	if err != nil {
		panic(err)
	}

	//Read one line at a time
	csvReader := csv.NewReader(fileReader)
	csvReader.LazyQuotes = true
	isFirstRow := true
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		//First row is just the headers, we will skip it
		if isFirstRow {
			isFirstRow = false
			continue
		}

		cleanURL := ParseURL(record[0])
		fmt.Println(cleanURL)
	}

	//store in Key-value store with value 1 or increment existing occurrence

}

//ParseURL parses a thing
//Parse out just the Method and path (case insensitive)
func ParseURL(url string) string {
	parts := strings.SplitAfter(url, "?")
	urlSansQuestionMark := strings.TrimSuffix(parts[0], "?")
	urlSansSlash := strings.TrimSuffix(urlSansQuestionMark, "/")
	return urlSansSlash
}
