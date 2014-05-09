package gitmob

import (
	"fmt"
	"os"
	"strings"

	"github.com/fraenkel/candiedyaml"
)

type Configuration struct {
	RawAuthors map[string]string `yaml:"authors"`
	RawEmail map[string]string `yaml:"email"`
	RawEmailAddresses map[string]string `yaml:"email_addresses"`
}

func LoadConfiguration(path string) Configuration {
	file, err := os.Open(path)

	if err != nil {
		println("File does not exist:", err.Error())
		os.Exit(1)
	}

	config := Configuration{}
	decoder := candiedyaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		println("Failed to decode document:", err.Error())
	}

	return config
}

func (config Configuration) Authors() Authors {
	authors := make(Authors)

	for initials, author := range config.RawAuthors {
		parts := strings.Split(string(author), ";")
		name := parts[0]
		email := config.RawEmailAddresses[initials]
		if "" == email {
			domain := config.RawEmail["domain"]
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

	return authors

}
