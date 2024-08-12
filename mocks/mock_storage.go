// Code generated by MockGen. DO NOT EDIT.
// Source: mishin-shortener/internal/app/handlers (interfaces: AbstractStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAbstractStorage is a mock of AbstractStorage interface.
type MockAbstractStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAbstractStorageMockRecorder
}

// MockAbstractStorageMockRecorder is the mock recorder for MockAbstractStorage.
type MockAbstractStorageMockRecorder struct {
	mock *MockAbstractStorage
}

// NewMockAbstractStorage creates a new mock instance.
func NewMockAbstractStorage(ctrl *gomock.Controller) *MockAbstractStorage {
	mock := &MockAbstractStorage{ctrl: ctrl}
	mock.recorder = &MockAbstractStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAbstractStorage) EXPECT() *MockAbstractStorageMockRecorder {
	return m.recorder
}

// Finish mocks base method.
func (m *MockAbstractStorage) Finish() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finish")
	ret0, _ := ret[0].(error)
	return ret0
}

// Finish indicates an expected call of Finish.
func (mr *MockAbstractStorageMockRecorder) Finish() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockAbstractStorage)(nil).Finish))
}

// Get mocks base method.
func (m *MockAbstractStorage) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAbstractStorageMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAbstractStorage)(nil).Get), arg0)
}

// Ping mocks base method.
func (m *MockAbstractStorage) Ping(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockAbstractStorageMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockAbstractStorage)(nil).Ping), arg0)
}

// Push mocks base method.
func (m *MockAbstractStorage) Push(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockAbstractStorageMockRecorder) Push(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockAbstractStorage)(nil).Push), arg0, arg1)
}

// PushBatch mocks base method.
func (m *MockAbstractStorage) PushBatch(arg0 *map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushBatch indicates an expected call of PushBatch.
func (mr *MockAbstractStorageMockRecorder) PushBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushBatch", reflect.TypeOf((*MockAbstractStorage)(nil).PushBatch), arg0)
}
