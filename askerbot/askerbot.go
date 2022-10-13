package askerbot

import (
	"fmt"
	"io"
)

type QuestionAsker interface {
	AskQuestion(user, question string) error
}

type TermAsker interface {
	AskInput(question string, isSkippable bool) (error, string, bool)
}

type DataProvider interface {
	ProvideUser() (error, string)
	ProvideQuestions() []string
	ProvideRepetitions() (error, int)
}

func Run(writer io.Writer, asker QuestionAsker, dataProvider DataProvider) {
	fmt.Fprintln(writer, "Welcome to NGL Asker BOT! o/")

	err, user := dataProvider.ProvideUser()
	if err != nil {
		return
	}

	questions := dataProvider.ProvideQuestions()
	if len(questions) == 0 {
		return
	}

	err, repeat := dataProvider.ProvideRepetitions()
	if err != nil {
		return
	}

	output := `

----------------------------
Started to ask questions
User %s

`

	fmt.Fprintf(writer, output, user)

	for _, question := range questions {
		for i := 1; i <= repeat; i++ {
			err := asker.AskQuestion(user, question)
			if err == nil {
				fmt.Fprintln(writer, "Asked question: ", question)
			}
		}
	}
}
