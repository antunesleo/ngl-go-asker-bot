package termasker

import (
	"bufio"
	"errors"
	"log"
)

type TermAsker struct {
	Scanner *bufio.Scanner
}

func (ta TermAsker) AskInput(question string, isSkippable bool) (error, string, bool) {
	if isSkippable {
		question = question + " | [s] to skip"
	}
	log.Println(question)
	var result string
	if ta.Scanner.Scan() {
		result = ta.Scanner.Text()
	} else {
		log.Fatalln("Error")
		return errors.New("Failed to scan inputs"), "", false
	}
	if result == "s" {
		return nil, "", true
	}
	if result == "" {
		log.Fatalln("This information is required, aborting")
		return errors.New("User didn't type any value"), "", false
	}
	return nil, result, false
}
