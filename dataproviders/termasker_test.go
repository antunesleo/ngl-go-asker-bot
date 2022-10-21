package dataproviders

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

	assertOutput(output, "The earth is flat?", t)
	assertNotError(err, t)
	assertInputYes(input, t)
	assertNotSkipped(skipped, t)
}

func TestShouldSkipQuestion(t *testing.T) {
	scanner := getScanner("s")

	var output bytes.Buffer

	ta := TermAsker{Writer: &output, Scanner: scanner}
	err, input, skipped := ta.AskInput("The earth is flat?", true)

	assertOutput(output, "The earth is flat? | [s] to skip", t)
	assertNotError(err, t)
	assertEmpty(input, t)
	assertSkipped(skipped, t)
}

func assertOutput(output bytes.Buffer, want_output string, t *testing.T) {
	if !strings.Contains(output.String(), want_output) {
		t.Errorf("Expected output to be The earth is flat? | [s] to skip, got %v instead", output.String())
	}
}

func assertSkipped(skipped bool, t *testing.T) {
	if !skipped {
		t.Errorf("Exected skipped to be true, got false instead")
	}
}

func assertEmpty(input string, t *testing.T) {
	if input != "" {
		t.Errorf("Expected input to be s, got %v instead", input)
	}
}

func assertInputYes(input string, t *testing.T) {
	if input != "yes" {
		t.Errorf("Expected input to be yes, got %v instead", input)
	}
}

func assertNotSkipped(skipped bool, t *testing.T) {
	if skipped {
		t.Errorf("Exected skipped to be false, got true instead")
	}
}
