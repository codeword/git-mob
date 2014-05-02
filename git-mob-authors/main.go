package main

import (
	"os"
	"fmt"

	"github.com/hiremaga/git-mob"
)

func main() {
	initialsList := os.Args[1:]
	file, err := os.Open(fmt.Sprintf("%s/.git-authors", os.Getenv("HOME")))
	if err != nil {
		println("File does not exist:", err.Error())
		os.Exit(1)
	}

	config := gitmob.NewConfiguration(file)
	authors := config.Authors()

	for _, initials := range initialsList {
		fmt.Printf("%#v\n", authors[initials])
	}
}
