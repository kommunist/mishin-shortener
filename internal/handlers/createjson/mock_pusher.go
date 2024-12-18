// Code generated by MockGen. DO NOT EDIT.
// Source: mishin-shortener/internal/handlers/createjson (interfaces: Pusher)

// Package createjson is a generated GoMock package.
package createjson

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPusher is a mock of Pusher interface.
type MockPusher struct {
	ctrl     *gomock.Controller
	recorder *MockPusherMockRecorder
}

// MockPusherMockRecorder is the mock recorder for MockPusher.
type MockPusherMockRecorder struct {
	mock *MockPusher
}

// NewMockPusher creates a new mock instance.
func NewMockPusher(ctrl *gomock.Controller) *MockPusher {
	mock := &MockPusher{ctrl: ctrl}
	mock.recorder = &MockPusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPusher) EXPECT() *MockPusherMockRecorder {
	return m.recorder
}

// Push mocks base method.
func (m *MockPusher) Push(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockPusherMockRecorder) Push(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockPusher)(nil).Push), arg0, arg1, arg2, arg3)
}
