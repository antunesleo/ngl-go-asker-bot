package termasker

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func getScanner(answer string) *bufio.Scanner {
	var reader io.Reader = strings.NewReader(answer)
	scanner := bufio.NewScanner(reader)
	return scanner
}

func TestShouldAskInput(t *testing.T) {
	scanner := getScanner("yes")

	var output bytes.Buffer

	ta := TermAsker{Writer: &output, Scanner: scanner}
	err, input, skipped := ta.AskInput("The earth is flat?", false)

	if !strings.Contains(output.String(), "The earth is flat?") {
		t.Errorf("Expected output to be The earth is flat?, got %v instead", output.String())
	}
	if err != nil {
		t.Errorf("Expected error to be nil")
	}
	if input != "yes" {
		t.Errorf("Expected input to be yes, got %v instead", input)
	}
	if skipped {
		t.Errorf("Exected skipped to be false, got true instead")
	}
}

func TestShouldSkipQuestion(t *testing.T) {
	scanner := getScanner("s")

	var output bytes.Buffer

	ta := TermAsker{Writer: &output, Scanner: scanner}
	err, input, skipped := ta.AskInput("The earth is flat?", true)

	if !strings.Contains(output.String(), "The earth is flat? | [s] to skip") {
		t.Errorf("Expected output to be The earth is flat? | [s] to skip, got %v instead", output.String())
	}
	if err != nil {
		t.Errorf("Expected error to be nil")
	}
	if input != "" {
		t.Errorf("Expected input to be s, got %v instead", input)
	}
	if !skipped {
		t.Errorf("Exected skipped to be true, got false instead")
	}
}
