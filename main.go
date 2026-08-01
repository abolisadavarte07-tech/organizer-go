package main

import (
	"fmt"

	"github.com/abolisadavarte07-tech/organizer-go/organizer"
)

func main() {
	fmt.Println(organizer.ExtensionMap[".jpg"])
	fmt.Println(organizer.ExtensionMap[".pdf"])
	fmt.Println(organizer.ExtensionMap[".mp3"])
	fmt.Println(organizer.ExtensionMap[".py"])
}