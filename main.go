package main

import (
	"fmt"
	"log"

	"github.com/abolisadavarte07-tech/organizer-go/organizer"
)

func main() {

	files, err := organizer.ScanDirectory("./TestFiles")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		category, found := organizer.GetCategory(file.Extension)

		if found {
			fmt.Printf("%-20s -> %s\n", file.Name, category)
		} else {
			fmt.Printf("%-20s -> Unsupported\n", file.Name)
		}
	}
}