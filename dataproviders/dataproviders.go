package dataproviders

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
)

type InputAsker interface {
	AskInput(question string, isSkippable bool) (error, string, bool)
}

type TermDataProvider struct {
	TermAsker InputAsker
}

func (tdp *TermDataProvider) ProvideUser() (error, string) {
	err, user, _ := tdp.TermAsker.AskInput("Type NGL user", false)
	return err, user
}

func (tdp *TermDataProvider) ProvideQuestions() []string {
	questions := []string{}
	for {
		err, question, skipped := tdp.TermAsker.AskInput("Type a question", true)
		if err != nil {
			return questions
		}
		if skipped {
			break
		}
		questions = append(questions, question)
	}
	return questions
}

var ErrInputCast = errors.New("Failed to cast input")

func (tdp *TermDataProvider) ProvideRepetitions() (error, int) {
	err, repetitionsStr, _ := tdp.TermAsker.AskInput("How many times should repeat questions?", false)
	repetitions, err := strconv.Atoi(repetitionsStr)
	if err != nil {
		return ErrInputCast, 0
	}
	return err, repetitions
}

type JsonFileDataProvider struct {
	Path string
}

type Data struct {
	User string `json:user`
}

var ErrFailedToReadJsonFile = errors.New("Failed to read json config")

func (jfdp *JsonFileDataProvider) ProvideUser() (error, string) {
	jsonFile, err := os.Open(jfdp.Path + "/" + "data.json")
	defer jsonFile.Close()

	if err != nil {
		return ErrFailedToReadJsonFile, ""
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data Data
	json.Unmarshal(byteValue, &data)
	return nil, data.User
}
