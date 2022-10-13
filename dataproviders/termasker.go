package dataproviders

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
)

type TermAsker struct {
	Scanner *bufio.Scanner
	Writer  io.Writer
}

func New(scanner *bufio.Scanner, writer io.Writer) *TermAsker {
	return &TermAsker{Scanner: scanner, Writer: writer}
}

func (ta TermAsker) AskInput(question string, isSkippable bool) (error, string, bool) {
	if isSkippable {
		question = question + " | [s] to skip"
	}
	fmt.Fprintln(ta.Writer, question)
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
		fmt.Fprintln(ta.Writer, "This information is required, aborting")
		return errors.New("User didn't type any value"), "", false
	}
	return nil, result, false
}
