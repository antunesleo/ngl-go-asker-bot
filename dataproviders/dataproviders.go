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
	data *Data
}

type Data struct {
	User        string   `json:"user"`
	Questions   []string `json:"questions"`
	Repetitions int      `json:"repetitions"`
}

var ErrFailedToReadJsonFile = errors.New("Failed to read json config")

func (jfdp *JsonFileDataProvider) ProvideUser() (error, string) {
	err, data := jfdp.getData()
	if err != nil {
		return err, ""
	}
	return nil, data.User
}

func (jfdp *JsonFileDataProvider) getData() (error, *Data) {
	if jfdp.data != nil {
		return nil, jfdp.data
	}
	filePath := jfdp.Path + "/" + "data.json"
	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()

	if err != nil {
		return ErrFailedToReadJsonFile, &Data{}
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data *Data
	json.Unmarshal(byteValue, &data)
	jfdp.data = data
	return nil, data
}

func (jfdp *JsonFileDataProvider) ProvideQuestions() (error, []string) {
	err, data := jfdp.getData()
	if err != nil {
		return err, []string{}
	}
	return nil, data.Questions
}

func (jfdp *JsonFileDataProvider) ProvideRepetitions() (error, int) {
	err, data := jfdp.getData()
	if err != nil {
		return err, 0
	}
	return nil, data.Repetitions
}
