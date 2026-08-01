package main

import (
	"fmt"
	"log"

	"github.com/abolisadavarte07-tech/organizer-go/organizer"
)

func main() {

	files, err := organizer.ScanDirectory(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("%s -> %s\n", file.Name, file.Extension)
	}
}