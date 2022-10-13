// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/antunesleo/ngl-go-asker-bot/askerbot (interfaces: QuestionAsker,DataProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQuestionAsker is a mock of QuestionAsker interface.
type MockQuestionAsker struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionAskerMockRecorder
}

// MockQuestionAskerMockRecorder is the mock recorder for MockQuestionAsker.
type MockQuestionAskerMockRecorder struct {
	mock *MockQuestionAsker
}

// NewMockQuestionAsker creates a new mock instance.
func NewMockQuestionAsker(ctrl *gomock.Controller) *MockQuestionAsker {
	mock := &MockQuestionAsker{ctrl: ctrl}
	mock.recorder = &MockQuestionAskerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuestionAsker) EXPECT() *MockQuestionAskerMockRecorder {
	return m.recorder
}

// AskQuestion mocks base method.
func (m *MockQuestionAsker) AskQuestion(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskQuestion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AskQuestion indicates an expected call of AskQuestion.
func (mr *MockQuestionAskerMockRecorder) AskQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskQuestion", reflect.TypeOf((*MockQuestionAsker)(nil).AskQuestion), arg0, arg1)
}

// MockDataProvider is a mock of DataProvider interface.
type MockDataProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDataProviderMockRecorder
}

// MockDataProviderMockRecorder is the mock recorder for MockDataProvider.
type MockDataProviderMockRecorder struct {
	mock *MockDataProvider
}

// NewMockDataProvider creates a new mock instance.
func NewMockDataProvider(ctrl *gomock.Controller) *MockDataProvider {
	mock := &MockDataProvider{ctrl: ctrl}
	mock.recorder = &MockDataProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataProvider) EXPECT() *MockDataProviderMockRecorder {
	return m.recorder
}

// ProvideQuestions mocks base method.
func (m *MockDataProvider) ProvideQuestions() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideQuestions")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ProvideQuestions indicates an expected call of ProvideQuestions.
func (mr *MockDataProviderMockRecorder) ProvideQuestions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideQuestions", reflect.TypeOf((*MockDataProvider)(nil).ProvideQuestions))
}

// ProvideRepetitions mocks base method.
func (m *MockDataProvider) ProvideRepetitions() (error, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideRepetitions")
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// ProvideRepetitions indicates an expected call of ProvideRepetitions.
func (mr *MockDataProviderMockRecorder) ProvideRepetitions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideRepetitions", reflect.TypeOf((*MockDataProvider)(nil).ProvideRepetitions))
}

// ProvideUser mocks base method.
func (m *MockDataProvider) ProvideUser() (error, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideUser")
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// ProvideUser indicates an expected call of ProvideUser.
func (mr *MockDataProviderMockRecorder) ProvideUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideUser", reflect.TypeOf((*MockDataProvider)(nil).ProvideUser))
}
