package main

import (
	"fmt"
	"os"

	"github.com/hiremaga/git-mob"
)

func main() {
	initialsList := os.Args[1:]
	path := fmt.Sprintf("%s/.git-authors", os.Getenv("HOME"))
	config := gitmob.LoadConfiguration(path)
	authors := config.Authors()

	for _, initials := range initialsList {
		fmt.Printf("%#v\n", authors[initials])
	}
}
