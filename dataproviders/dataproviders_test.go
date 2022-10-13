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
		err, got := tdp.ProvideQuestions()
		assertNotError(err, t)
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
		AssertError(ErrInputCast, err, t)
	})
}

func TestConfigFileDataProvider(t *testing.T) {
	t.Run("should provide user", func(t *testing.T) {
		want := "breno"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		jfdp := JsonFileDataProvider{Path: "./testdata"}
		err, got := jfdp.ProvideUser()
		assertNotError(err, t)
		assertEqual(got, want, t)
	})

	t.Run("should not provide user when no config file", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		jfdp := JsonFileDataProvider{Path: "./wrongdir"}
		err, _ := jfdp.ProvideUser()
		AssertError(err, ErrFailedToReadJsonFile, t)
	})

	t.Run("should provide questions", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		jfdp := JsonFileDataProvider{Path: "./testdata"}
		_, questions := jfdp.ProvideQuestions()
		assertIntEqual(len(questions), 2, t)
		assertEqual("Is the earth flat?", questions[0], t)
		assertEqual("Is the flat earth?", questions[1], t)
	})

	t.Run("should not provide questions when no config file", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		jfdp := JsonFileDataProvider{Path: "./wrongdir"}
		err, _ := jfdp.ProvideQuestions()
		AssertError(err, ErrFailedToReadJsonFile, t)
	})

	t.Run("should provide repetitions", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		jfdp := JsonFileDataProvider{Path: "./testdata"}
		_, repetitions := jfdp.ProvideRepetitions()
		assertIntEqual(repetitions, 100, t)
	})

	t.Run("should get cached data", func(t *testing.T) {
		want := "breno"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		jfdp := JsonFileDataProvider{Path: "./testdata"}
		err, data := jfdp.getData()
		assertNotError(err, t)
		assertEqual(data.User, want, t)

		jfdp.Path = "./wrongpath"
		err, data = jfdp.getData()
		assertNotError(err, t)
		assertEqual(data.User, want, t)
	})
}

func TestCreateDataProvider(t *testing.T) {
	t.Run("create json file data provider when has json file", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockInputAsker := mocks.NewMockInputAsker(mockCtrl)

		dataProvider := CreateDataProvider("./testdata", mockInputAsker)
		_, ok := dataProvider.(*JsonFileDataProvider)
		assertTrue(ok, t)
	})

	t.Run("create term data provider when has not json file", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockInputAsker := mocks.NewMockInputAsker(mockCtrl)
		dataProvider := CreateDataProvider("./wrongpath", mockInputAsker)
		_, ok := dataProvider.(*TermDataProvider)
		assertTrue(ok, t)

	})
}

func AssertError(want error, got error, t *testing.T) {
	if want != got {
		t.Fatalf("expect error to be %v got %v", want, got)
	}
}

func assertFalse(value bool, t *testing.T) {
	if value {
		t.Errorf("Expected it to be True")
	}
}

func assertTrue(value bool, t *testing.T) {
	if !value {
		t.Errorf("Expected it to be True")
	}
}

func assertEqual(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("Expected %s, got %s", want, got)
	}
}

func assertIntEqual(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func assertNotError(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("expected error to be nil")
	}
}
