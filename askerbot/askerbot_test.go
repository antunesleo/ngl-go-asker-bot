package askerbot

import (
	"bytes"
	"testing"

	"github.com/antunesleo/ngl-go-asker-bot/mocks"
	"github.com/golang/mock/gomock"
)

func TestShouldAskQuestionsOnNGL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := "breno"
	question := "The flat is earth?"
	questions := []string{question}

	mockQuestionAsker := mocks.NewMockQuestionAsker(mockCtrl)
	mockDataProvider := mocks.NewMockDataProvider(mockCtrl)

	mockQuestionAsker.EXPECT().AskQuestion(user, question).Return(nil).Times(2)
	mockDataProvider.EXPECT().ProvideUser().Return(nil, user)
	mockDataProvider.EXPECT().ProvideQuestions().Return(questions)
	mockDataProvider.EXPECT().ProvideRepetitions().Return(nil, 2)

	Run(&bytes.Buffer{}, mockQuestionAsker, mockDataProvider)
}
