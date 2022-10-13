package dataproviders

import (
	"errors"
	"testing"

	"github.com/antunesleo/ngl-go-asker-bot/mocks"
	"github.com/golang/mock/gomock"
)

func TestTermDataProvider(t *testing.T) {
	t.Run("should provide user", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		want := "breno"
		mockInputAsker := mocks.NewMockInputAsker(mockCtrl)
		first := mockInputAsker.EXPECT().AskInput("Type NGL user", false).Return(nil, want, false)

		gomock.InOrder(
			first,
		)

		tdp := TermDataProvider{mockInputAsker}
		err, got := tdp.ProvideUser()
		assertNotError(err, t)
		assertEqual(got, want, t)
	})

	t.Run("should provide questions", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		want := []string{"Is the earth flat?"}
		mockInputAsker := mocks.NewMockInputAsker(mockCtrl)
		first := mockInputAsker.EXPECT().AskInput("Type a question", true).Return(nil, want[0], false)
		second := mockInputAsker.EXPECT().AskInput("Type a question", true).Return(nil, "s", true)

		gomock.InOrder(
			first,
			second,
		)

		tdp := TermDataProvider{mockInputAsker}
		got := tdp.ProvideQuestions()
		assertIntEqual(len(got), 1, t)
		assertEqual(got[0], want[0], t)
	})

	t.Run("should provide repetitions", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		want := 2
		mockInputAsker := mocks.NewMockInputAsker(mockCtrl)
		mockInputAsker.EXPECT().AskInput("How many times should repeat questions?", false).Return(nil, "2", false)

		tdp := TermDataProvider{mockInputAsker}
		err, got := tdp.ProvideRepetitions()
		assertNotError(err, t)
		assertIntEqual(got, want, t)
	})

	t.Run("provide repetitions should return error when fail to cast input", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockInputAsker := mocks.NewMockInputAsker(mockCtrl)
		mockInputAsker.EXPECT().AskInput("How many times should repeat questions?", false).Return(errors.New(""), "", false)

		tdp := TermDataProvider{mockInputAsker}
		err, _ := tdp.ProvideRepetitions()
		AssertError(err, t)
	})
}

func TestConfigFileDataProvider(t *testing.T) {
	t.Run("should provide user", func(t *testing.T) {
		want := "breno"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		cfdp := JsonFileDataProvider{Path: "/home/leo/projects/ngl-go-asker-bot/mocks"}
		err, got := cfdp.ProvideUser()
		assertNotError(err, t)
		assertEqual(got, want, t)
	})
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
