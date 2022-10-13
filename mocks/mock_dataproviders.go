// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/antunesleo/ngl-go-asker-bot/dataproviders (interfaces: InputAsker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInputAsker is a mock of InputAsker interface.
type MockInputAsker struct {
	ctrl     *gomock.Controller
	recorder *MockInputAskerMockRecorder
}

// MockInputAskerMockRecorder is the mock recorder for MockInputAsker.
type MockInputAskerMockRecorder struct {
	mock *MockInputAsker
}

// NewMockInputAsker creates a new mock instance.
func NewMockInputAsker(ctrl *gomock.Controller) *MockInputAsker {
	mock := &MockInputAsker{ctrl: ctrl}
	mock.recorder = &MockInputAskerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInputAsker) EXPECT() *MockInputAskerMockRecorder {
	return m.recorder
}

// AskInput mocks base method.
func (m *MockInputAsker) AskInput(arg0 string, arg1 bool) (error, string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskInput", arg0, arg1)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// AskInput indicates an expected call of AskInput.
func (mr *MockInputAskerMockRecorder) AskInput(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskInput", reflect.TypeOf((*MockInputAsker)(nil).AskInput), arg0, arg1)
}