package askerbot

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/antunesleo/ngl-go-asker-bot/mocks"
	"github.com/golang/mock/gomock"
)

func TestShouldAskQuestionsOnNGL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := "breno"
	question := "The flat is earth?"

	mockQuestionAsker := mocks.NewMockQuestionAsker(mockCtrl)
	mockTermAsker := mocks.NewMockTermAsker(mockCtrl)

	mockQuestionAsker.EXPECT().AskQuestion(user, question).Return(nil).Times(1)
	first := mockTermAsker.EXPECT().AskInput("Type NGL user", false).Return(nil, user, false)
	second := mockTermAsker.EXPECT().AskInput("Type a question", true).Return(nil, question, false)
	third := mockTermAsker.EXPECT().AskInput("Type a question", true).Return(nil, "s", true)

	gomock.InOrder(
		first,
		second,
		third,
	)

	fmt.Println("mockQuestionAsker", mockQuestionAsker)
	fmt.Println("mockTermAsker", mockTermAsker)

	Run(&bytes.Buffer{}, mockQuestionAsker, mockTermAsker)
}
