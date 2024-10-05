package hangman

import (
	"io/ioutil"
	"log"
	"strings"
)

// OpenJose : Open hangman.txt file and return a string sliced fore each positions of the hangman
func OpenJose() []string {
	JoseFile, err := ioutil.ReadFile("Extra/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	JoseStr := string(JoseFile)                 // Transform our file into string
	JoseSlice := strings.Split(JoseStr, "\n\n") 
	return JoseSlice
}
