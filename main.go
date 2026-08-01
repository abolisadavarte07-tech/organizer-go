package main

import (
	"log"

	"github.com/abolisadavarte07-tech/organizer-go/organizer"
)

func main() {

	err := organizer.Organize("./TestFiles")

	if err != nil {
		log.Fatal(err)
	}
}