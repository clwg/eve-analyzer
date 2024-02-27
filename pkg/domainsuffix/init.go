package domainsuffix

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	publicSuffixes []string
	initOnce       sync.Once
)

func init() {
	initOnce.Do(loadPublicSuffixList)
}

func loadPublicSuffixList() {

	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the location of the public_suffix_list.dat file from the environment variable
	publicSuffixListLocation := os.Getenv("PUBLIC_SUFFIX_LIST")

	file, err := os.Open(publicSuffixListLocation)
	if err != nil {
		panic("failed to load public suffix list: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore comments and empty lines
		if line == "" || line[0] == '/' || line[0] == ' ' {
			continue
		}
		publicSuffixes = append(publicSuffixes, line)
	}
}
