package askerbot

import (
	"fmt"
	"io"
)

type QuestionAsker interface {
	AskQuestion(user, question string) error
}

type DataProvider interface {
	ProvideUser() (error, string)
	ProvideQuestions() (error, []string)
	ProvideRepetitions() (error, int)
}

func Run(writer io.Writer, asker QuestionAsker, dataProvider DataProvider) {
	fmt.Fprintln(writer, "Welcome to NGL Asker BOT! o/")

	err, user := dataProvider.ProvideUser()
	if err != nil {
		line := fmt.Sprintf("Skipping, Failed to get user data, err %v", err)
		fmt.Fprintln(writer, line)
		return
	}

	err, questions := dataProvider.ProvideQuestions()
	if err != nil {
		line := fmt.Sprintf("Skipping, Failed to get questions data, err %v", err)
		fmt.Fprintln(writer, line)
		return
	}
	if len(questions) == 0 {
		return
	}

	err, repeat := dataProvider.ProvideRepetitions()
	if err != nil {
		line := fmt.Sprintf("Skipping, Failed to get repeat data, err %v", err)
		fmt.Fprintln(writer, line)
		return
	}

	askQuestions(writer, user, questions, repeat, asker)
}

func askQuestions(writer io.Writer, user string, questions []string, repeat int, asker QuestionAsker) {
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
			} else {
				line := fmt.Sprintf("Failed to ask question %s, err %v", question, err)
				fmt.Fprintln(writer, line)
			}
		}
	}
}
