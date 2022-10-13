package dataproviders

import (
	"errors"
	"strconv"
)

type TermAsker interface {
	AskInput(question string, isSkippable bool) (error, string, bool)
}

type TermDataProvider struct {
	termAsker TermAsker
}

func (tdp *TermDataProvider) ProvideUser() (error, string) {
	err, user, _ := tdp.termAsker.AskInput("Type NGL user", false)
	return err, user
}

func (tdp *TermDataProvider) ProvideQuestion() (error, string, bool) {
	err, user, skipped := tdp.termAsker.AskInput("Type a question", true)
	return err, user, skipped
}

var ErrInputCast = errors.New("Failed to cast input")

func (tdp *TermDataProvider) ProvideRepetitions() (error, int) {
	err, repetitionsStr, _ := tdp.termAsker.AskInput("How many times should repeat questions?", false)
	repetitions, err := strconv.Atoi(repetitionsStr)
	if err != nil {
		return ErrInputCast, 0
	}
	return err, repetitions
}

type JSONFileDataProvider struct{}
