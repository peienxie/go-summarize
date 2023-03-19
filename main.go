package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var client Client

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing OPENAI_API_KEY environment variable")
	}
	client = NewClient(apiKey)
}

func main() {
	// Parse command line arguments
	var text string
	var filePath string
	flag.Usage = func() {
		w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
		fmt.Fprintf(w, "Usage of %s:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.StringVar(&text, "text", "", "Text to summarize")
	flag.StringVar(&filePath, "file", "", "Path to file containing text to summarize")
	flag.Parse()

	// Make sure one input method is used
	if text == "" && filePath == "" {
		fmt.Println("Invalid arguments. Please specify either -text or -file")
		flag.Usage()
		os.Exit(1)
	} else if text != "" && filePath != "" {
		fmt.Println("Only one input method allowed: -text or -file")
		flag.Usage()
		os.Exit(1)
	}

	// Read input text from file or argument
	var input string
	if filePath != "" {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal("Failed to read file:", err)
		}
		input = string(data)
	} else {
		input = text
	}

	// Use the OPENAI client and request summary
	response, err := client.SummarizeText(context.Background(), input)
	if err != nil {
		log.Fatal("Failed to generate summary:", err)
	}

	// Print summary
	fmt.Println("Summary:", response)
}
