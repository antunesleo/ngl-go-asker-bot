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

func Run(writer io.Writer, asker QuestionAsker, termAsker TermAsker) {
	questions := []string{}

	fmt.Fprintln(writer, "Welcome to NGL Asker BOT! o/")
	err, user, _ := termAsker.AskInput("Type NGL user", false)
	if err != nil {
		return
	}

	for {
		err, question, skipped := termAsker.AskInput("Type a question", true)
		if err != nil {
			return
		}
		if skipped {
			break
		}
		questions = append(questions, question)
	}

	output := `

----------------------------
Started to ask questions
User %s

`

	fmt.Fprintf(writer, output, user)

	for _, question := range questions {
		err := asker.AskQuestion(user, question)
		if err == nil {
			fmt.Fprintln(writer, "Asked question: ", question)
		}
	}
}
