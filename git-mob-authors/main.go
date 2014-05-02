package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/fraenkel/candiedyaml"
)

type configuration struct {
	Authors map[string]string `yaml:"authors"`
	Email map[string]string `yaml:"email"`
	EmailAddresses map[string]string `yaml:"email_addresses"`
}

type authors map[string]Author


type Author struct {
	Name string
	Email string
}

func main() {
	file, err := os.Open(fmt.Sprintf("%s/.git-authors", os.Getenv("HOME")))
	if err != nil {
		println("File does not exist:", err.Error())
		os.Exit(1)
	}

	config := configuration{}
	decoder := candiedyaml.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		println("Failed to decode document:", err.Error())
	}

	authors := make(map[string]Author)

	for initials, author := range config.Authors {
		parts := strings.Split(string(author), ";")
		name := parts[0]
		email := config.EmailAddresses[initials]
		if "" == email {
			domain := config.Email["domain"]
			var username string
			if len(parts) < 2 {
				username = strings.Split(name, " ")[0]
			} else {
				username = parts[1]
			}
			email = fmt.Sprintf("%s@%s", username, domain)
		}
		authors[initials] = Author{Name: name, Email: email}
	}

	for _, initials := range os.Args[1:] {
		fmt.Printf("%#v\n", authors[initials])
	}
}
