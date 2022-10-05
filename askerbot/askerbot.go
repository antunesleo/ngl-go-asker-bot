package askerbot

import (
	"log"
)

type QuestionAsker interface {
	AskQuestion(user, question string) error
}

type TermAsker interface {
	AskInput(question string, isSkippable bool) (error, string, bool)
}

func Run(asker QuestionAsker, termAsker TermAsker) {
	questions := []string{}

	log.Println("Welcome to NGL Asker BOT! o/")
	err, user, _ := termAsker.AskInput("Type NGL user", false)
	if err != nil {
		return
	}

	addQuestion := true
	for addQuestion {
		err, question, skipped := termAsker.AskInput("Type a question", true)
		if err != nil {
			return
		}
		if skipped {
			addQuestion = false
			break
		}
		questions = append(questions, question)
	}

	log.Println("")
	log.Println("")
	log.Println("-----------------")
	log.Println("Started to ask questions")
	log.Println("User: ", user)

	for _, question := range questions {
		err := asker.AskQuestion(user, question)
		if err == nil {
			log.Println("Asked question: ", question)
		}
	}
}
