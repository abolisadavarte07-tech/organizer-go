package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abolisadavarte07-tech/organizer-go/organizer"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
        fmt.Println("go run . all")
        fmt.Println("go run . safe")
        fmt.Println("go run . image")
        fmt.Println("go run . audio")
        fmt.Println("go run . video")
        fmt.Println("go run . text")
        fmt.Println("go run . vector")
        fmt.Println("go run . gif")
        fmt.Println("go run . photoshop")
        fmt.Println("go run . pdf")
        fmt.Println("go run . python")
        fmt.Println("go run . font")
		fmt.Println("go run . office")
		return
	}

	command := os.Args[1]

	count, err := organizer.CountMovableFiles("./TestFiles", command)
    if err != nil {
	    log.Fatal(err)
    }

    if count == 0 {
	    fmt.Println("No matching files found.")
	    return
    }

    fmt.Printf("ATTENTION: %d file(s) will be moved.\n", count)
    fmt.Print("Continue? (y/n): ")

    var choice string
    fmt.Scanln(&choice)

    if choice != "y" && choice != "Y" {
	    fmt.Println("Operation cancelled.")
	    return
    }

	err = organizer.Organize("./TestFiles", command)

	if err != nil {
		log.Fatal(err)
	}
}