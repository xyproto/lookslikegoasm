package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xyproto/lookslikegoasm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <assembly_source_file>")
		os.Exit(1)
	}
	filename := os.Args[1]
	sourceBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filename, err)
	}
	if lookslikegoasm.Consider(string(sourceBytes)) {
		fmt.Println("This looks like Go/Plan9 Assembly")
	} else {
		fmt.Println("This does not look like Go/Plan9 Assembly")
	}
}
