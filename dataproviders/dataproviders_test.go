package dataproviders

import (
	"errors"
	"testing"

	"github.com/antunesleo/ngl-go-asker-bot/mocks"
	"github.com/golang/mock/gomock"
)

func TestShouldProvideUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	want := "breno"
	mockTermAsker := mocks.NewMockTermAsker(mockCtrl)
	first := mockTermAsker.EXPECT().AskInput("Type NGL user", false).Return(nil, want, false)

	gomock.InOrder(
		first,
	)

	tdp := TermDataProvider{mockTermAsker}
	err, got := tdp.ProvideUser()
	assertNotError(err, t)
	assertEqual(got, want, t)
}

func TestShouldProvideQuestion(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	want := "Is the earth flat?"
	mockTermAsker := mocks.NewMockTermAsker(mockCtrl)
	mockTermAsker.EXPECT().AskInput("Type a question", true).Return(nil, want, false)

	tdp := TermDataProvider{mockTermAsker}
	err, got, skipped := tdp.ProvideQuestion()
	assertNotError(err, t)
	assertEqual(got, want, t)
	assertFalse(skipped, t)
}

func TestShouldProvideRepetitions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	want := 2
	mockTermAsker := mocks.NewMockTermAsker(mockCtrl)
	mockTermAsker.EXPECT().AskInput("How many times should repeat questions?", false).Return(nil, "2", false)

	tdp := TermDataProvider{mockTermAsker}
	err, got := tdp.ProvideRepetitions()
	assertNotError(err, t)
	assertIntEqual(got, want, t)
}

func TestProvideRepetitionsShouldReturnErrorWhenFailToCastInput(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTermAsker := mocks.NewMockTermAsker(mockCtrl)
	mockTermAsker.EXPECT().AskInput("How many times should repeat questions?", false).Return(errors.New(""), "", false)

	tdp := TermDataProvider{mockTermAsker}
	err, _ := tdp.ProvideRepetitions()
	AssertError(err, t)
}

func AssertError(err error, t *testing.T) {
	if err != ErrInputCast {
		t.Fatalf("expect error to be %v got %v", ErrInputCast, err)
	}
}

func assertFalse(skipped bool, t *testing.T) {
	if skipped {
		t.Errorf("Expected it to be True")
	}
}

func assertEqual(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("Expected user to be %s, got %s", want, got)
	}
}

func assertIntEqual(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("Expected user to be %v, got %v", want, got)
	}
}

func assertNotError(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("expected error to be nil")
	}
}
